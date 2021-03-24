package main

import (
	"html/template"
	"net/http"
	"path"
)

func renderView(w http.ResponseWriter, templateName string, data map[string]interface{}, componentTemplateNames ...string) {
	var componentPaths []string
	for _, componentName := range componentTemplateNames {
		componentPath := path.Join("views", componentName+".html")
		componentPaths = append(componentPaths, componentPath)
	}

	mainTemplatePath := path.Join("views", templateName+".html")
	componentPaths = append(componentPaths, mainTemplatePath)

	t := template.Must(template.ParseFiles(componentPaths...))
	err := t.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/" {
			templateComponents := []string{"_header", "_footer"}
			renderView(w, "404", nil, templateComponents...)
		} else {
			templateComponents := []string{"_header", "_navbar", "_sidebar", "_footer"}
			renderView(w, "dashboard", nil, templateComponents...)
		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
