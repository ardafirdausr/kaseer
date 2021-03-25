package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

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
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/products" {
			data := map[string]interface{}{
				"Templates": []string{"_meta", "_navbar", "_sidebar", "_script"},
			}
			renderView(w, "404", data)
		} else {
			data := map[string]interface{}{
				"Templates":  []string{"_meta", "_navbar", "_sidebar", "_script"},
				"Title":      "Dashboard",
				"ActiveMenu": "dashboard",
			}
			renderView(w, "dashboard", data)
		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case (r.Method == http.MethodGet) && (r.URL.Path == "/products"):
	case (r.Method == http.MethodPost) && (r.URL.Path == "/products"):
		{
			product := Product{}
			products, err := product.GetAllProducts()
			if err != nil {
				log.Println(err.Error())
				data := map[string]interface{}{
					"Templates": []string{"_meta", "_script"},
				}
				renderView(w, "500", data)
			}

			data := map[string]interface{}{
				"Templates":  []string{"_meta", "_navbar", "_sidebar", "_script"},
				"Title":      "All Products",
				"ActiveMenu": "products",
				"Products":   products,
			}
			renderView(w, "products", data)
		}
	case (r.Method == http.MethodGet) && (r.URL.Path == "/products/create"):
		{
			data := map[string]interface{}{
				"Templates":  []string{"_meta", "_navbar", "_sidebar", "_script"},
				"Title":      "Create Product",
				"ActiveMenu": "products",
			}
			renderView(w, "products", data)
		}
	case (r.Method == http.MethodGet) && (r.URL.Path == "/products/{id}/delete"):
	case (r.Method == http.MethodGet) && (r.URL.Path == "/products/{id}/edit"):
	case (r.Method == http.MethodPut) && (r.URL.Path == "/products/{id}"):
		{

		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
