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

	web.Static("/static", "web/assets")

	userController := controller.NewUserController(app.Usecases)

	// guest routes
	authGuestRouter := web.Group("/auth", middleware.SessionGuest())
	authGuestRouter.GET("/login", userController.ShowLoginForm)
	authGuestRouter.POST("/login", userController.Login)

	// authenticated rotues
	authAuthenticatedUserRouter := web.Group("/auth", middleware.SessionAuth())
	authAuthenticatedUserRouter.POST("/logout", userController.Logout)

	authenticatedGroup := web.Group("", middleware.SessionAuth())

	// Profile Routes
	profileRouter := authenticatedGroup.Group("/profile")
	profileRouter.GET("", userController.ShowUserProfile)
	profileRouter.GET("/edit/password", userController.ShowEditUserPasswordForm)
	profileRouter.GET("/edit", userController.ShowEditUserProfileForm)
	profileRouter.POST("", userController.UpdateUserProfile)
	profileRouter.POST("/password", userController.UpdateUserPassword)

	// Order Routes
	orderController := controller.NewOrderController(app.Usecases)
	orderRouter := authenticatedGroup.Group("/orders")
	orderRouter.GET("/create", orderController.ShowCreateOrderForm)
	orderRouter.GET("/total", orderController.GetTotalOrdersData)
	orderRouter.GET("/latest-income", orderController.GetLatestIncomeData)
	orderRouter.GET("/annual-income", orderController.GetAnnualIncomeData)
	orderRouter.GET("/:orderId", orderController.GetOrderDetailData)
	orderRouter.GET("", orderController.ShowAllOrders)
	orderRouter.POST("", orderController.CreateOrder)

	// Product Routes
	productController := controller.NewProductController(app.Usecases)
	productRouter := authenticatedGroup.Group("/products")
	productRouter.GET("/create", productController.ShowCreateProductForm)
	productRouter.GET("/bestseller", productController.GetBestSellerProductsData)
	productRouter.GET("/:productId/edit", productController.ShowEditProductForm)
	productRouter.GET("", productController.ShowAllProducts)
	productRouter.POST("/:productId/update", productController.UpdateProduct)
	productRouter.POST("/:productId/delete", productController.DeleteProduct)
	productRouter.POST("", productController.CreateProduct)

	// Dashboard route
	dashboardController := controller.NewDashboardController(app.Usecases)
	authenticatedGroup.GET("/dashboard", dashboardController.ShowDashboard)

	// Redirect Route
	web.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/dashboard")
	})

	server.Start(web)
}
