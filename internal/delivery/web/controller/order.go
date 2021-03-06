package controller

import (
	"net/http"
	"strconv"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/app"
	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderUc   internal.OrderUsecase
	productUc internal.ProductUsecase
}

func NewOrderController(ucs *app.Usecases) *OrderController {
	orderUc := ucs.OrderUsecase
	productUc := ucs.ProductUsecase
	return &OrderController{orderUc, productUc}
}

func (oc OrderController) ShowAllOrders(c echo.Context) error {
	ctx := c.Request().Context()
	orders, err := oc.orderUc.GetAllOrders(ctx)
	if err != nil {
		return err
	}

	data := echo.Map{"Orders": orders}
	return renderPage(c, "orders", "All Orders", data)
}

func (oc OrderController) ShowCreateOrderForm(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := oc.productUc.GetAllProducts(ctx)
	if err != nil {
		return err
	}

	data := echo.Map{"Products": products}
	return renderPage(c, "order_create", "Create New Order", data)
}

func (oc OrderController) GetTotalOrdersData(c echo.Context) error {
	ctx := c.Request().Context()
	orderType := c.FormValue("type")
	var totalOrderCount int
	switch orderType {
	case "day":
		dailyOrder, err := oc.orderUc.GetDailyOrderCount(ctx)
		if err != nil {
			return echo.ErrInternalServerError
		}

		totalOrderCount = dailyOrder
	default:
		totalOrder, err := oc.orderUc.GetTotalOrderCount(ctx)
		if err != nil {
			return echo.ErrInternalServerError
		}

		totalOrderCount = totalOrder
	}

	return responseJson(c, http.StatusOK, "Success", totalOrderCount)
}

func (oc OrderController) GetLatestIncomeData(c echo.Context) error {
	ctx := c.Request().Context()
	orderType := c.FormValue("type")
	var totalIncome int
	switch orderType {
	case "day":
		dailyIncome, err := oc.orderUc.GetLastDayIncome(ctx)
		if err != nil {
			return echo.ErrInternalServerError
		}
		totalIncome = dailyIncome
	case "month":
		monthlyIncome, err := oc.orderUc.GetLastMonthIncome(ctx)
		if err != nil {
			return echo.ErrInternalServerError
		}

		totalIncome = monthlyIncome
	}

	return responseJson(c, http.StatusOK, "Success", totalIncome)
}

func (oc OrderController) GetAnnualIncomeData(c echo.Context) error {
	ctx := c.Request().Context()
	annualIncomes, err := oc.orderUc.GetAnnualIncome(ctx)
	if err != nil {
		return err
	}

	return responseJson(c, http.StatusOK, "Success", annualIncomes)
}

func (oc OrderController) GetOrderDetailData(c echo.Context) error {
	paramOrderID := c.Param("orderId")
	orderID, err := strconv.ParseInt(paramOrderID, 10, 64)
	if err != nil {
		return echo.ErrInternalServerError
	}

	ctx := c.Request().Context()
	orderItems, err := oc.orderUc.GetOrderItems(ctx, orderID)
	if err != nil {
		return err
	}

	return responseJson(c, http.StatusOK, "Success", orderItems)
}

func (oc OrderController) CreateOrder(c echo.Context) error {
	var orderParam entity.CreateOrderParam
	if err := c.Bind(&orderParam); err != nil {
		return responseJson(c, http.StatusInternalServerError, "Failed processing data", nil)
	}

	err := c.Validate(&orderParam)
	if ev, ok := err.(entity.ErrValidation); ok {
		return responseJson(c, http.StatusBadRequest, "Invalid data", ev.Errors)
	}

	if err != nil {
		return responseJson(c, http.StatusBadRequest, "Invalid data", nil)
	}

	ctx := c.Request().Context()
	order, err := oc.orderUc.Create(ctx, orderParam)
	if ev, ok := err.(entity.ErrValidation); ok {
		return responseErrorJson(c, http.StatusBadRequest, ev.Message, ev.Errors)
	}

	if err != nil {
		return responseJson(c, http.StatusInternalServerError, "Failed creating data", nil)
	}

	return responseJson(c, http.StatusCreated, "Success creating order", order)
}
