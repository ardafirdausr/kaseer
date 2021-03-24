package controllers

import (
	"html/template"
	"net/http"
	"path"
)

func renderView(w http.ResponseWriter, filename string, data map[string]interface{}) {
	t := template.Must(template.ParseFiles(
		path.Join("views", "_header.html"),
		path.Join("views", "_footer.html"),
		path.Join("views", filename+".html"),
	))

	err := t.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
