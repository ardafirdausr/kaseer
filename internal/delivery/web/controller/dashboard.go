package controller

import (
	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/app"
	"github.com/labstack/echo/v4"
)

type DashboardController struct {
	productUc internal.ProductUsecase
	orderUc   internal.OrderUsecase
}

func NewDashboardController(ucs *app.Usecases) *DashboardController {
	return &DashboardController{
		productUc: ucs.ProductUsecase,
		orderUc:   ucs.OrderUsecase,
	}
}

func (dc DashboardController) ShowDashboard(c echo.Context) error {
	return renderPage(c, "dashboard", "Dashboard", nil)
}
