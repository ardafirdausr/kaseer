package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/products", ProductsHandler)
	mux.HandleFunc("/", IndexHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "80"
	}
	fmt.Println("Running app on port " + port)
	http.ListenAndServe(":"+port, mux)
}
