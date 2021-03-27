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
	data := M{
		"Templates":    []string{"_meta", "_script"},
		"Title":        "Login",
		"ErrorMessage": GoPosSession.Flashes("error_message"),
	}
	renderView(w, "login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
		return
	}

	email := r.Form.Get("email")
	user, err := findUserByEmail(email)
	if err != nil || user == nil {
		GoPosSession.AddFlash("Invalid Email Or Password", "error_message")
		GoPosSession.Save(r, w)
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	password := r.Form.Get("password")
	isEqual := user.CheckPassword(password)
	if !isEqual {
		GoPosSession.AddFlash("Invalid Email Or Password", "error_message")
		GoPosSession.Save(r, w)
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	GoPosSession.Values["user"] = user
	GoPosSession.Save(r, w)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	GoPosSession.Values["user"] = nil
	GoPosSession.Save(r, w)
	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func ShowDashboard(w http.ResponseWriter, r *http.Request) {
	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Dashboard",
		"ActiveMenu": "dashboard",
	}
	renderView(w, "dashboard", data)
}

func ShowAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := GetAllProducts()
	if err != nil {
		log.Println(err.Error())
		data := M{
			"Templates": []string{"_meta", "_script"},
		}
		renderView(w, "500", data)
		return
	}

	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "All Products",
		"ActiveMenu": "products",
		"Products":   products,
	}
	renderView(w, "products", data)
}

func ShowCreateProductForm(w http.ResponseWriter, r *http.Request) {
	data := M{
		"Templates":    []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":        "Create Product",
		"ActiveMenu":   "products",
		"ErrorMessage": GoPosSession.Flashes("error_message"),
	}
	renderView(w, "product_create", data)
}

func ShowEditProductForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, _ := strconv.Atoi(vars["productId"])
	product, err := FindProductById(productId)
	if err != nil || product == nil {
		renderErrorPage(w, http.StatusNotFound)
	}

	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Edit Product",
		"ActiveMenu": "products",
		"Product":    product,
	}
	renderView(w, "product_edit", data)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
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

		GoPosSession.AddFlash("Invalid data", "error_message")

		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
		return
	}

	err = product.Save()
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
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
		renderErrorPage(w, http.StatusNotFound)
	}

	err = r.ParseForm()
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
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
		renderErrorPage(w, http.StatusNotFound)
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
		renderView(w, "500", data)
		return
	}

	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "All Products",
		"ActiveMenu": "orders",
		"Products":   orders,
	}
	renderView(w, "orders", data)
}

func ShowCreateOrderForm(w http.ResponseWriter, r *http.Request) {
	products, err := GetAllProducts()
	if err != nil {
		log.Println(err.Error())
		data := M{
			"Templates": []string{"_meta", "_script"},
		}
		renderView(w, "500", data)
		return
	}

	data := M{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Create Order",
		"ActiveMenu": "orderss",
		"Products":   products,
	}
	renderView(w, "order_create", data)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
		return
	}

	now := time.Now()

	payload := struct {
		OrderItems []OrderItem `json:"order_items"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError)
		return
	}

	order := &Order{}
	order.Code = fmt.Sprintf("%d%d%d", now.Year(), now.Month(), now.Day())

	// order.Stock, _ = strconv.Atoi(r.Form.Get("stock"))
	// order.Price, _ = strconv.Atoi(r.Form.Get("price"))

	validate := validator.New()
	err = validate.Struct(order)
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
	}

	err = order.Save()
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		http.Redirect(w, r, "/products/create", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/order_create", http.StatusSeeOther)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderErrorPage(w, http.StatusNotFound)
}

func renderErrorPage(w http.ResponseWriter, errorCode int) {
	templateName := strconv.Itoa(errorCode)
	data := M{
		"Templates": []string{"_meta", "_script"},
	}
	renderView(w, templateName, data)
}

func renderView(w http.ResponseWriter, templateName string, data M) {
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
