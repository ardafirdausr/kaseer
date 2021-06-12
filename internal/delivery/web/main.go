package web

import (
	"net/http"

	"github.com/ardafirdausr/go-pos/internal/app"
	"github.com/ardafirdausr/go-pos/internal/delivery/web/controller"
	"github.com/ardafirdausr/go-pos/internal/delivery/web/middleware"
	"github.com/ardafirdausr/go-pos/internal/delivery/web/server"
	"github.com/labstack/echo/v4"
)

func Start(app *app.App) {
	web := server.New()

	web.Static("/static/", "web/assets")

	userController := controller.NewUserController(app.Usecases)

	authGuestRouter := web.Group("/auth", middleware.SessionGuest())
	authGuestRouter.GET("/login", userController.ShowLoginForm)
	authGuestRouter.POST("/login", userController.Login)

	authAuthenticatedUserRouter := web.Group("/auth", middleware.SessionAuth())
	authAuthenticatedUserRouter.POST("/logout", userController.Logout)

	authenticatedGroup := web.Group("", middleware.SessionAuth())

	// profileRouter := router.PathPrefix("/profile").Subrouter()
	// profileRouter.Use(AuthMiddleware)
	// profileRouter.HandleFunc("", ShowUserProfile).Methods("GET")
	// profileRouter.HandleFunc("/edit/password", showEditUserPasswordForm).Methods("GET")
	// profileRouter.HandleFunc("/edit", showEditUserProfileForm).Methods("GET")
	// profileRouter.HandleFunc("", UpdateUserProfile).Methods("POST")
	// profileRouter.HandleFunc("/password", UpdateUserPassword).Methods("POST")

	// orderRouter := router.PathPrefix("/orders").Subrouter()
	// orderRouter.Use(AuthMiddleware)
	orderController := controller.NewOrderController(app.Usecases)
	orderRouter := authenticatedGroup.Group("/orders")
	// orderRouter.GET("/create", orderController.ShowCreateOrderForm)
	orderRouter.GET("/total", orderController.GetTotalOrdersData)
	orderRouter.GET("/latest-income", orderController.GetLatestIncomeData)
	orderRouter.GET("/annual-income", orderController.GetAnnualIncomeData)
	// orderRouter.HandleFunc("/{orderId:[0-9]+}", GetOrderDetailData).Methods("GET")
	// orderRouter.HandleFunc("", ShowAllOrders).Methods("GET")
	// orderRouter.HandleFunc("", CreateOrder).Methods("POST")

	productController := controller.NewProductController(app.Usecases)
	productRouter := authenticatedGroup.Group("/products")
	productRouter.GET("/bestseller", productController.GetBestSellerProductsData)
	// productRouter.HandleFunc("/create", ShowCreateProductForm).Methods("GET")
	// productRouter.HandleFunc("/{productId:[0-9]+}/edit", ShowEditProductForm).Methods("GET")
	// productRouter.HandleFunc("/{productId:[0-9]+}/update", UpdateProduct).Methods("POST")
	// productRouter.HandleFunc("/{productId:[0-9]+}/delete", DeleteProduct).Methods("POST")
	// productRouter.HandleFunc("", ShowAllProducts).Methods("GET")
	// productRouter.HandleFunc("", CreateProduct).Methods("POST")

	dashboardController := controller.NewDashboardController(app.Usecases)
	authenticatedGroup.GET("/dashboard", dashboardController.ShowDashboard)

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	// })

	web.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/dashboard")
	})

	server.Start(web)
}
