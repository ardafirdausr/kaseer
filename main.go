package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	fmt.Println("Serving app on port 9000")
	http.ListenAndServe(":9000", mux)
}
