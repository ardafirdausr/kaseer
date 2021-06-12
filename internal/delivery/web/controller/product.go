package controller

import (
	"net/http"

	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/app"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUc internal.ProductUsecase
}

func NewProductController(ucs *app.Usecases) *ProductController {
	productUc := ucs.ProductUsecase
	return &ProductController{productUc}
}

func (pc ProductController) GetBestSellerProductsData(c echo.Context) error {
	products, err := pc.productUc.GetBestSellerProducts()
	if err != nil {
		return err
	}

	return json(c, http.StatusOK, "Success", products)
}
