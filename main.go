package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var DB *sql.DB

var SessionStore *sessions.CookieStore

const SESSIONNAME = "go_pos_session"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file\n%v\n", err.Error())
	}

	DB, err = ConnectToDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"))
	if err != nil {
		log.Fatalf("Failed to connect to the database\n%v\n", err.Error())
	}
	defer DB.Close()

	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		sessionKey = "secret-session-key"
	}
	SessionStore = NewSessionStore(sessionKey)

	router := InitRouter()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Running app on port " + port)
	http.ListenAndServe(":"+port, router)
}
