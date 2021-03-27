package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SessionName)
	data := M{
		"Templates":    []string{"_meta", "_script"},
		"Title":        "Login",
		"ErrorMessage": session.Flashes("error_message"),
	}
	renderView(w, r, "login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SessionName)

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
	err = session.Save(r, w)
	if err != nil {
		log.Println(err.Error())
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SessionName)
	session.Values["user"] = nil
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
	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Edit Profile",
		"ActiveMenu": "",
	}
	renderView(w, r, "profile_edit", data)
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(4096); err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

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

func ShowCreateProductForm(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SessionName)
	data := M{
		"Templates":    []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":        "Create Product",
		"ActiveMenu":   "products",
		"ErrorMessage": session.Flashes("error_message"),
	}
	renderView(w, r, "product_create", data)
}

func ShowEditProductForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, _ := strconv.Atoi(vars["productId"])
	product, err := FindProductById(productId)
	if err != nil || product == nil {
		renderErrorPage(w, r, http.StatusNotFound)
	}

	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Edit Product",
		"ActiveMenu": "products",
		"Product":    product,
	}
	renderView(w, r, "product_edit", data)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SessionName)

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
	err = validate.Struct(product)
	if err != nil {
		log.Println(err.Error())

		session.AddFlash(err.Error(), "error_message")

		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
		return
	}

	err = product.Save()
	if err != nil {
		log.Println(err.Error())
		session.AddFlash(err.Error(), "error_message")

		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, _ := strconv.Atoi(vars["productId"])
	product, err := FindProductById(productId)
	if err != nil || product == nil {
		renderErrorPage(w, r, http.StatusNotFound)
	}

	err = r.ParseForm()
	if err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	product.Code = r.Form.Get("code")
	product.Name = r.Form.Get("name")
	product.Stock, _ = strconv.Atoi(r.Form.Get("stock"))
	product.Price, _ = strconv.Atoi(r.Form.Get("price"))
	err = product.Update()
	if err != nil {
		editUrl := fmt.Sprintf("/products/%d/edit", productId)
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

	err = product.Delete()
	if err != nil {
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
		"Products":   orders,
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
		"ActiveMenu": "orderss",
		"Products":   products,
	}
	renderView(w, r, "order_create", data)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	now := time.Now()

	payload := struct {
		OrderItems []OrderItem `json:"order_items"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		renderErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	order := &Order{}
	order.Code = fmt.Sprintf("%d%d%d", now.Year(), now.Month(), now.Day())

	// order.Stock, _ = strconv.Atoi(r.Form.Get("stock"))
	// order.Price, _ = strconv.Atoi(r.Form.Get("price"))

	validate := validator.New()
	err = validate.Struct(order)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
	}

	err = order.Save()
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/order_create", http.StatusSeeOther)
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
	session, _ := SessionStore.Get(r, SessionName)
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
