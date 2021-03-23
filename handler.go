package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
