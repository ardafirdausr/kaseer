package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/ardafirdausr/kaseer/internal/mocks"
	"github.com/stretchr/testify/assert"
)

var products = []*entity.Product{
	{
		ID:    1,
		Code:  "prod-1",
		Name:  "prod 1",
		Price: 5000,
		Stock: 100,
	}, {
		ID:    2,
		Code:  "prod-2",
		Name:  "prod 2",
		Price: 10000,
		Stock: 200,
	},
}

var productSales = []*entity.ProductSale{
	{
		ID:   1,
		Code: "prod-1",
		Name: "prod 1",
		Sale: 20,
	}, {
		ID:   2,
		Code: "prod-2",
		Name: "prod 2",
		Sale: 50,
	},
}

func Test_GetAllProducts_Failed(t *testing.T) {
	ctx := context.TODO()
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetAllProducts", ctx).Return(nil, errors.New("failed get products"))

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.GetAllProducts(ctx)
	assert.NotNil(t, err)
	assert.Nil(t, aProducts)
}

func Test_GetAllProducts_Success(t *testing.T) {
	ctx := context.TODO()
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetAllProducts", ctx).Return(products, nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.GetAllProducts(ctx)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(t, aProducts)
}

func Test_GetProductByID_Failed(t *testing.T) {
	ctx := context.TODO()
	expectedProduct := products[0]
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByID", ctx, expectedProduct.ID).Return(nil, errors.New("failed get product by id"))

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.GetProductByID(ctx, expectedProduct.ID)
	assert.NotNil(t, err)
	assert.Nil(t, aProducts)
}

func Test_GetProductByID_Success(t *testing.T) {
	ctx := context.TODO()
	expectedProduct := products[0]
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByID", ctx, expectedProduct.ID).Return(expectedProduct, nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.GetProductByID(ctx, expectedProduct.ID)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(expectedProduct, aProducts)
}

func Test_GetProductByCode_Failed(t *testing.T) {
	ctx := context.TODO()
	expectedProduct := products[0]
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByCode", ctx, expectedProduct.Code).Return(nil, errors.New("failed get product by code"))

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.GetProductByCode(ctx, expectedProduct.Code)
	assert.NotNil(t, err)
	assert.Nil(t, aProducts)
}

func Test_GetProductByCode_Success(t *testing.T) {
	ctx := context.TODO()
	expectedProduct := products[0]
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByCode", ctx, expectedProduct.Code).Return(expectedProduct, nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.GetProductByCode(ctx, expectedProduct.Code)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(expectedProduct, aProducts)
}

func Test_GetBestSellerProducts_Failed(t *testing.T) {
	ctx := context.TODO()
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetBestSellerProducts", ctx).Return(nil, errors.New("failed get product sales"))

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.GetBestSellerProducts(ctx)
	assert.NotNil(t, err)
	assert.Nil(t, aProducts)
}

func Test_GetBestSellerProducts_Success(t *testing.T) {
	ctx := context.TODO()
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetBestSellerProducts", ctx).Return(productSales, nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.GetBestSellerProducts(ctx)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(productSales, aProducts)
}

func Test_CreateProduct_Failed_When_Product_Code_Already_Exists(t *testing.T) {
	ctx := context.TODO()
	existProduct := products[0]
	createParam := entity.CreateProductParam{
		Code:  existProduct.Code,
		Name:  existProduct.Name,
		Price: existProduct.Price,
		Stock: existProduct.Stock,
	}

	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByCode", ctx, existProduct.Code).Return(existProduct, nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.CreateProduct(ctx, createParam)
	assert.Nil(t, aProducts)
	assert.NotNil(t, err)
	assert.IsType(t, entity.ErrItemAlreadyExists{}, err)
}

func Test_CreateProduct_Failed_When_Creating_Product(t *testing.T) {
	ctx := context.TODO()
	createParam := entity.CreateProductParam{
		Code:  "new-product",
		Name:  "New Product",
		Price: 10000,
		Stock: 50,
	}

	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByCode", ctx, createParam.Code).Return(nil, nil)
	mockProductRepo.On("Create", ctx, createParam).Return(nil, errors.New("failed create product"))

	productUsecase := NewProductUsecase(mockProductRepo)
	aProducts, err := productUsecase.CreateProduct(ctx, createParam)
	assert.Nil(t, aProducts)
	assert.NotNil(t, err)
}

func Test_CreateProduct_Success(t *testing.T) {
	ctx := context.TODO()
	createParam := entity.CreateProductParam{
		Code:  "new-product",
		Name:  "New Product",
		Price: 10000,
		Stock: 50,
	}
	eProduct := &entity.Product{
		ID:    99,
		Code:  "new-product",
		Name:  "New Product",
		Price: 10000,
		Stock: 50,
	}

	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByCode", ctx, createParam.Code).Return(nil, nil)
	mockProductRepo.On("Create", ctx, createParam).Return(eProduct, nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	aProduct, err := productUsecase.CreateProduct(ctx, createParam)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eProduct, aProduct)
}

func Test_UpdateProduct_Failed_When_Product_Code_Already_Exists(t *testing.T) {
	ctx := context.TODO()
	var productID int64 = 2
	updateParam := entity.UpdateProductParam{
		Code:  "updated-product",
		Name:  "Updated Product",
		Price: 15000,
		Stock: 89,
	}

	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByCode", ctx, updateParam.Code).Return(products[0], nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	isUpdated, err := productUsecase.UpdateProduct(ctx, productID, updateParam)
	assert.False(t, isUpdated)
	assert.NotNil(t, err)
	assert.IsType(t, entity.ErrItemAlreadyExists{}, err)
}

func Test_UpdateProduct_Failed_WhenUpdating_Product(t *testing.T) {
	ctx := context.TODO()
	var productID int64 = 1
	updateParam := entity.UpdateProductParam{
		Code:  "updated-product",
		Name:  "Updated Product",
		Price: 15000,
		Stock: 89,
	}

	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByCode", ctx, updateParam.Code).Return(nil, nil)
	mockProductRepo.On("UpdateByID", ctx, productID, updateParam).Return(false, errors.New("failed update product"))

	productUsecase := NewProductUsecase(mockProductRepo)
	isUpdated, err := productUsecase.UpdateProduct(ctx, productID, updateParam)
	assert.NotNil(t, err)
	assert.False(t, isUpdated)
}

func Test_UpdateProduct_Success(t *testing.T) {
	ctx := context.TODO()
	var productID int64 = 1
	updateParam := entity.UpdateProductParam{
		Code:  "updated-product",
		Name:  "Updated Product",
		Price: 15000,
		Stock: 89,
	}

	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductByCode", ctx, updateParam.Code).Return(nil, nil)
	mockProductRepo.On("UpdateByID", ctx, productID, updateParam).Return(true, nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	isUpdated, err := productUsecase.UpdateProduct(ctx, productID, updateParam)
	assert.Nil(t, err)
	assert.True(t, isUpdated)
}

func Test_DeleteProduct_Failed(t *testing.T) {
	ctx := context.TODO()
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("DeleteByID", ctx, products[0].ID).Return(false, errors.New("failed to delete product"))

	productUsecase := NewProductUsecase(mockProductRepo)
	isDeleted, err := productUsecase.DeleteProduct(ctx, products[0].ID)
	assert.NotNil(t, err)
	assert.False(t, isDeleted)
}

func Test_DeleteProduct_Success(t *testing.T) {
	ctx := context.TODO()
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("DeleteByID", ctx, products[0].ID).Return(true, nil)

	productUsecase := NewProductUsecase(mockProductRepo)
	isDeleted, err := productUsecase.DeleteProduct(ctx, products[0].ID)
	assert.Nil(t, err)
	assert.True(t, isDeleted)
}
