package controller

import (
	"net/http"

	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/app"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderUc internal.OrderUsecase
}

func NewOrderController(ucs *app.Usecases) *OrderController {
	orderUc := ucs.OrderUsecase
	return &OrderController{orderUc}
}

func (oc OrderController) ShowCreateOrderForm(c echo.Context) error {
	return echo.ErrNotFound
}

func (oc OrderController) GetTotalOrdersData(c echo.Context) error {
	orderType := c.FormValue("type")
	var totalOrderCount int
	switch orderType {
	case "day":
		dailyOrder, err := oc.orderUc.GetDailyOrderCount()
		if err != nil {
			return echo.ErrInternalServerError
		}

		totalOrderCount = dailyOrder
	default:
		totalOrder, err := oc.orderUc.GetTotalOrderCount()
		if err != nil {
			return echo.ErrInternalServerError
		}

		totalOrderCount = totalOrder
	}

	return json(c, http.StatusOK, "Success", totalOrderCount)
}

func (oc OrderController) GetLatestIncomeData(c echo.Context) error {
	orderType := c.FormValue("type")
	var totalIncome int
	switch orderType {
	case "day":
		dailyIncome, err := oc.orderUc.GetLastDayIncome()
		if err != nil {
			return echo.ErrInternalServerError
		}
		totalIncome = dailyIncome
	case "month":
		monthlyIncome, err := oc.orderUc.GetLastMonthIncome()
		if err != nil {
			return echo.ErrInternalServerError
		}

		totalIncome = monthlyIncome
	}

	return json(c, http.StatusOK, "Success", totalIncome)
}

func (oc OrderController) GetAnnualIncomeData(c echo.Context) error {
	annualIncomes, err := oc.orderUc.GetAnnualIncome()
	if err != nil {
		return err
	}

	return json(c, http.StatusOK, "Success", annualIncomes)
}
