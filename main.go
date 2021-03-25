package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file\n%v\n", err.Error())
	}

	DB, err := ConnectToDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"))
	if err != nil {
		log.Fatalf("Failed to connect to the database\n%v\n", err.Error())
	}
	defer DB.Close()

	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	// router.Use(mux.CORSMethodMiddleware(router))

	// router.StrictSlash(true)
	// router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	// 	w.WriteHeader(http.StatusNoContent)
	// 	return
	// })

	fs := http.FileServer(http.Dir("assets"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("/create", ShowCreateProductForm).Methods("GET")
	productRouter.HandleFunc("/{productId:[0-9]+}/edit", ShowEditProductForm).Methods("GET")
	productRouter.HandleFunc("/{productId:[0-9]+}/update", UpdateProduct).Methods("POST")
	productRouter.HandleFunc("/{productId:[0-9]+}/delete", DeleteProduct).Methods("POST")
	productRouter.HandleFunc("", ShowAllProducts).Methods("GET")
	productRouter.HandleFunc("", CreateProduct).Methods("POST")

	router.HandleFunc("/", ShowDashboard).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "80"
	}
	fmt.Println("Running app on port " + port)
	http.ListenAndServe(":"+port, router)
}
