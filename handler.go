package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	errorMessage := session.Flashes("error_message")
	session.Save(r, w)

	data := M{
		"Templates":    []string{"_meta", "_script"},
		"Title":        "Login",
		"ErrorMessage": errorMessage,
	}
	renderView(w, r, "login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)

	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	email := r.Form.Get("email")
	user, err := findUserByEmail(email)
	if err != nil || user == nil {
		session.AddFlash("Invalid Email Or Password", "error_message")
		session.Save(r, w)
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	password := r.Form.Get("password")
	isEqual := user.CheckPassword(password)
	if !isEqual {
		session.AddFlash("Invalid Email Or Password", "error_message")
		session.Save(r, w)
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	session.Values["user_id"] = user.ID
	if err := session.Save(r, w); err != nil {
		log.Println(err.Error())
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Profile",
		"ActiveMenu": "",
	}
	renderView(w, r, "profile", data)
}

func showEditUserProfileForm(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	errorMessage := session.Flashes("error_message")
	session.Save(r, w)

	data := M{
		"Templates":    []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":        "Edit Profile",
		"ActiveMenu":   "",
		"ErrorMessage": errorMessage,
	}
	renderView(w, r, "profile_edit", data)
}

func showEditUserPasswordForm(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	errorMessage := session.Flashes("error_message")
	session.Save(r, w)

	data := M{
		"Templates":    []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":        "Edit Password",
		"ActiveMenu":   "",
		"ErrorMessage": errorMessage,
	}
	renderView(w, r, "profile_password", data)
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	user := session.Values["user"].(*User)

	if err := r.ParseForm(); err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	password := r.Form.Get("password")
	user.changePassword(password)
	if err := user.Update(); err != nil {
		session.AddFlash("Failed to update data", "error_message")
		session.Save(r, w)
		http.Redirect(w, r, "/profile/edit/password", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	user := session.Values["user"].(*User)

	if err := r.ParseMultipartForm(1024 * 5); err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	photoDirectory := filepath.Join("assets", "storage", "image")
	photoName := fmt.Sprintf("user-%d", user.ID)
	filename, err := SaveUploadedFile(r, "photo", photoDirectory, photoName)
	if err != nil {
		session.AddFlash(err.Error(), "error_message")
		session.Save(r, w)
		http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
		return
	}

	photoUrl := fmt.Sprintf("/static/storage/image/%s", filename)
	user.PhotoUrl = &photoUrl

	user.Name = r.Form.Get("name")
	user.Email = r.Form.Get("email")
	if err := user.Update(); err != nil {
		session.AddFlash("Failed to update data", "error_message")
		session.Save(r, w)
		http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
		return
	}

	session.Values["user"] = user
	session.Save(r, w)
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func ShowDashboard(w http.ResponseWriter, r *http.Request) {
	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Dashboard",
		"ActiveMenu": "dashboard",
	}
	renderView(w, r, "dashboard", data)
}

func ShowAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := GetAllProducts()
	if err != nil {
		log.Println(err.Error())
		data := M{
			"Templates": []string{"_meta", "_script"},
		}
		renderView(w, r, "500", data)
		return
	}

	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "All Products",
		"ActiveMenu": "products",
		"Products":   products,
	}
	renderView(w, r, "products", data)
}

func GetBestSellerProductsData(w http.ResponseWriter, r *http.Request) {
	products, err := GetBestSellerProducts()
	if err != nil {
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	data := M{"message": "success", "data": products}
	jsonResponse(w, data, http.StatusOK)
	return
}

func ShowCreateProductForm(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	errorMessage := session.Flashes("error_message")
	session.Save(r, w)

	data := M{
		"Templates":    []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":        "Create Product",
		"ActiveMenu":   "products",
		"ErrorMessage": errorMessage,
	}
	renderView(w, r, "product_create", data)
}

func ShowEditProductForm(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	errorMessage := session.Flashes("error_message")
	session.Save(r, w)

	vars := mux.Vars(r)
	productId, _ := strconv.Atoi(vars["productId"])
	product, err := FindProductById(productId)
	if err != nil || product == nil {
		renderErrorPage(w, r, http.StatusNotFound)
	}

	data := M{
		"Templates":    []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":        "Edit Product",
		"ActiveMenu":   "products",
		"Product":      product,
		"ErrorMessage": errorMessage,
	}
	renderView(w, r, "product_edit", data)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)

	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	product := &Product{}
	product.Code = r.Form.Get("code")
	product.Name = r.Form.Get("name")
	product.Stock, _ = strconv.Atoi(r.Form.Get("stock"))
	product.Price, _ = strconv.Atoi(r.Form.Get("price"))

	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		log.Println(err.Error())
		session.AddFlash(err.Error(), "error_message")
		session.Save(r, w)
		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
		return
	}

	if err := product.Save(); err != nil {
		log.Println(err.Error())
		session.AddFlash(err.Error(), "error_message")
		session.Save(r, w)
		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SESSIONNAME)

	vars := mux.Vars(r)
	productId, _ := strconv.Atoi(vars["productId"])
	product, err := FindProductById(productId)
	if err != nil || product == nil {
		renderErrorPage(w, r, http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	product.Code = r.Form.Get("code")
	product.Name = r.Form.Get("name")
	product.Stock, _ = strconv.Atoi(r.Form.Get("stock"))
	product.Price, _ = strconv.Atoi(r.Form.Get("price"))
	if err := product.Update(); err != nil {
		editUrl := fmt.Sprintf("/products/%d/edit", productId)
		session.AddFlash(err.Error(), "error_message")
		session.Save(r, w)
		http.Redirect(w, r, editUrl, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, _ := strconv.Atoi(vars["productId"])
	product, err := FindProductById(productId)
	if err != nil || product == nil {
		renderErrorPage(w, r, http.StatusNotFound)
	}

	if err := product.Delete(); err != nil {
		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func ShowAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := GetAllOrders()
	if err != nil {
		log.Println(err.Error())
		data := M{
			"Templates": []string{"_meta", "_script"},
		}
		renderView(w, r, "500", data)
		return
	}

	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "All Products",
		"ActiveMenu": "orders",
		"Orders":     orders,
	}
	renderView(w, r, "orders", data)
}

func ShowCreateOrderForm(w http.ResponseWriter, r *http.Request) {
	products, err := GetAllProducts()
	if err != nil {
		log.Println(err.Error())
		data := M{
			"Templates": []string{"_meta", "_script"},
		}
		renderView(w, r, "500", data)
		return
	}

	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Create Order",
		"ActiveMenu": "orders",
		"Products":   products,
	}
	renderView(w, r, "order_create", data)
}

func GetOrderDetailData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	orderId, _ := strconv.Atoi(vars["orderId"])
	products, err := GetOrderDetail(orderId)
	if err != nil {
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	data := M{"message": "success", "data": products}
	jsonResponse(w, data, http.StatusOK)
	return
}

func GetTotalOrdersData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	totalOrderType := r.Form.Get("type")
	val, err := GetTotalOrders(totalOrderType)
	if err != nil {
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	data := M{"message": "success", "data": val}
	jsonResponse(w, data, http.StatusOK)
	return
}

func GetLatestEarningData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	earningType := r.Form.Get("type")
	val, err := GetLastEarning(earningType)
	if err != nil {
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	data := M{"message": "success", "data": val}
	jsonResponse(w, data, http.StatusOK)
	return
}

func GetAnnualEarningData(w http.ResponseWriter, r *http.Request) {
	val, err := GetAnnualEarning()
	if err != nil {
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
	}

	data := M{"message": "success", "data": val}
	jsonResponse(w, data, http.StatusOK)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	order := &Order{}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(order); err != nil {
		log.Println(err.Error())
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(order); err != nil {
		log.Println(err.Error())
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	if err := order.Save(); err != nil {
		log.Println(err.Error())
		data := M{"error": err.Error()}
		jsonResponse(w, data, http.StatusInternalServerError)
		return
	}

	data := M{"message": "success"}
	jsonResponse(w, data, http.StatusCreated)
	return
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderErrorPage(w, r, http.StatusNotFound)
}

func renderErrorPage(w http.ResponseWriter, r *http.Request, errorCode int) {
	templateName := strconv.Itoa(errorCode)
	data := M{
		"Templates": []string{"_meta", "_script"},
	}
	renderView(w, r, templateName, data)
}

func renderView(w http.ResponseWriter, r *http.Request, templateName string, data M) {
	session, _ := SessionStore.Get(r, SESSIONNAME)
	data["User"] = session.Values["user"]

	var templatesPaths []string
	if templates, isExist := data["Templates"]; isExist {
		for _, template := range templates.([]string) {
			templatePath := path.Join("views", template+".html")
			templatesPaths = append(templatesPaths, templatePath)
		}
	}

	mainTemplatePath := path.Join("views", templateName+".html")
	templatesPaths = append(templatesPaths, mainTemplatePath)

	t := template.Must(template.ParseFiles(templatesPaths...))
	err := t.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func jsonResponse(w http.ResponseWriter, data M, status int) {
	var payload, err = json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}

	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	w.Write(payload)
}
