package main

import (
	"html/template"
	"net/http"
	"path"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			t := template.Must(template.ParseFiles(
				path.Join("views", "_header.html"),
				path.Join("views", "_footer.html"),
				path.Join("views", "index.html"),
			))

			data := map[string]interface{}{}
			data["Title"] = "Index"
			err := t.ExecuteTemplate(w, "index", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
