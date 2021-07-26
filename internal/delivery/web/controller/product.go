package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/app"
	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUc internal.ProductUsecase
}

func NewProductController(ucs *app.Usecases) *ProductController {
	productUc := ucs.ProductUsecase
	return &ProductController{productUc}
}

func (pc ProductController) ShowAllProducts(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := pc.productUc.GetAllProducts(ctx)
	if err != nil {
		return err
	}

	data := echo.Map{"Products": products}
	return renderPage(c, "products", "All Products", data)
}

func (pc ProductController) GetBestSellerProductsData(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := pc.productUc.GetBestSellerProducts(ctx)
	if err != nil {
		return err
	}

	return responseJson(c, http.StatusOK, "Success", products)
}

func (pc ProductController) ShowCreateProductForm(c echo.Context) error {
	return renderPage(c, "product_create", "Create Product", nil)
}

func (pc ProductController) ShowEditProductForm(c echo.Context) error {
	pid := c.Param("productId")
	productID, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		return echo.ErrNotFound
	}

	ctx := c.Request().Context()
	product, err := pc.productUc.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}

	data := echo.Map{"Product": product}
	return renderPage(c, "product_edit", "Edit Product", data)
}

func (pc ProductController) CreateProduct(c echo.Context) error {
	sess, _ := session.Get("kaseer", c)

	var param entity.CreateProductParam
	if err := c.Bind(&param); err != nil {
		return echo.ErrInternalServerError
	}

	err := c.Validate(&param)
	if ev, ok := err.(entity.ErrValidation); ok {
		sess.AddFlash(ev, "error_validation")
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			log.Println(err)
		}
		return c.Redirect(http.StatusSeeOther, "/products/create")
	}

	if err != nil {
		return echo.ErrInternalServerError
	}

	ctx := c.Request().Context()
	product, err := pc.productUc.CreateProduct(ctx, param)
	if eae, ok := err.(entity.ErrItemAlreadyExists); ok {
		msg := fmt.Sprintf("Failed creating product. %s", eae.Message)
		sess.AddFlash(msg, "error_message")
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusSeeOther, "/products/create")
	}

	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Success creating \"%s\"", product.Name)
	sess.AddFlash(msg, "success_message")
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/products")
}

func (pc ProductController) UpdateProduct(c echo.Context) error {
	sess, _ := session.Get("kaseer", c)

	pid := c.Param("productId")
	productID, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		return echo.ErrNotFound
	}

	ctx := c.Request().Context()
	_, err = pc.productUc.GetProductByID(ctx, productID)
	if _, ok := err.(entity.ErrNotFound); ok {
		return echo.ErrNotFound
	}

	var updateParam entity.UpdateProductParam
	if err := c.Bind(&updateParam); err != nil {
		return echo.ErrInternalServerError
	}

	err = c.Validate(&updateParam)
	if ev, ok := err.(entity.ErrValidation); ok {
		sess.AddFlash(ev, "error_validation")
		sess.Save(c.Request(), c.Response())
		editProductUrl := fmt.Sprintf("/products/%d/create", productID)
		return c.Redirect(http.StatusSeeOther, editProductUrl)
	}

	isUpdated, err := pc.productUc.UpdateProduct(ctx, productID, updateParam)
	if err != nil {
		return err
	}

	if !isUpdated {
		return echo.ErrInternalServerError
	}

	sess.AddFlash("Success Updating the Product", "success_message")
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/products")
}

func (pc ProductController) DeleteProduct(c echo.Context) error {
	pid := c.Param("productId")
	productID, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		return echo.ErrNotFound
	}

	ctx := c.Request().Context()
	isDeleted, err := pc.productUc.DeleteProduct(ctx, productID)
	if err != nil {
		return err
	}

	if !isDeleted {
		return echo.ErrInternalServerError
	}

	sess, _ := session.Get("kaseer", c)
	sess.AddFlash("Success Deleting Product", "success_message")
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/products")
}
