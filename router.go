package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {

	router := mux.NewRouter()

	router.Use(LoggingMiddleware)

	fs := http.FileServer(http.Dir("assets"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", ShowLoginForm).Methods("GET")
	authRouter.HandleFunc("/login", Login).Methods("POST")
	authRouter.HandleFunc("/logout", Logout).Methods("POST")

	profileRouter := router.PathPrefix("/profile").Subrouter()
	profileRouter.Use(AuthMiddleware)
	profileRouter.HandleFunc("", ShowUserProfile).Methods("GET")
	profileRouter.HandleFunc("/edit/password", showEditUserPasswordForm).Methods("GET")
	profileRouter.HandleFunc("/edit", showEditUserProfileForm).Methods("GET")
	profileRouter.HandleFunc("", UpdateUserProfile).Methods("POST")
	profileRouter.HandleFunc("/password", UpdateUserPassword).Methods("POST")

	orderRouter := router.PathPrefix("/orders").Subrouter()
	orderRouter.Use(AuthMiddleware)
	orderRouter.HandleFunc("/create", ShowCreateOrderForm).Methods("GET")
	orderRouter.HandleFunc("", ShowAllOrders).Methods("GET")
	orderRouter.HandleFunc("", CreateOrder).Methods("POST")

	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.Use(AuthMiddleware)
	productRouter.HandleFunc("/create", ShowCreateProductForm).Methods("GET")
	productRouter.HandleFunc("/{productId:[0-9]+}/edit", ShowEditProductForm).Methods("GET")
	productRouter.HandleFunc("/{productId:[0-9]+}/update", UpdateProduct).Methods("POST")
	productRouter.HandleFunc("/{productId:[0-9]+}/delete", DeleteProduct).Methods("POST")
	productRouter.HandleFunc("", ShowAllProducts).Methods("GET")
	productRouter.HandleFunc("", CreateProduct).Methods("POST")

	IndexRouter := router.PathPrefix("/dashboard").Subrouter()
	IndexRouter.Use(AuthMiddleware)
	IndexRouter.HandleFunc("", ShowDashboard).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})

	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	return router
}
