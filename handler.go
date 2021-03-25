package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

func ShowDashboard(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
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
		data := map[string]interface{}{
			"Templates": []string{"_meta", "_script"},
		}
		renderView(w, "500", data)
		return
	}

	data := map[string]interface{}{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "All Products",
		"ActiveMenu": "products",
		"Products":   products,
	}
	renderView(w, "products", data)
}

func ShowCreateProductForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Templates":  []string{"_meta", "_navbar", "_sidebar", "_footer", "_script"},
		"Title":      "Create Product",
		"ActiveMenu": "products",
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

	data := map[string]interface{}{
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

	// validate := validator.New()
	// err = validate.Struct(product)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	log.Println(err.Error())
	// 	http.Redirect(w, r, "/products/create", http.StatusSeeOther)
	// }

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

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderErrorPage(w, http.StatusNotFound)
}

func renderErrorPage(w http.ResponseWriter, errorCode int) {
	templateName := strconv.Itoa(errorCode)
	data := map[string]interface{}{
		"Templates": []string{"_meta", "_script"},
	}
	renderView(w, templateName, data)
}

func renderView(w http.ResponseWriter, templateName string, data map[string]interface{}) {
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
