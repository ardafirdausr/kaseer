package mysql

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllProducts_Failed_WhenSelectData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT * FROM products")
	mock.ExpectQuery(query).WillReturnError(errors.New("failed get products"))

	productRepository := NewProductRepository(db)
	products, err := productRepository.GetAllProducts(ctx)
	assert.NotNil(t, err)
	assert.Nil(t, products)
}

func Test_GetAllProducts_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var eProducts = sqlmock.
		NewRows([]string{"ID", "Code", "Name", "Price", "Stock", "CreatedAt", "UpdatedAt"}).
		AddRow(1, "prod-1", "Prod 1", 10000, 100, time.Now(), time.Now())
	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT * FROM products")
	mock.ExpectQuery(query).WillReturnRows(eProducts)

	productRepository := NewProductRepository(db)
	aProducts, err := productRepository.GetAllProducts(ctx)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eProducts, aProducts)
}

func Test_GetBestSellerProducts_Failed_WhenSelectData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta(`
		SELECT p.ID, p.Code, p.Name, SUM(oi.quantity) as total_sales
			FROM products AS p  JOIN order_items AS oi
			ON p.id = oi.product_id
			GROUP BY oi.product_id
			ORDER BY total_sales DESC
			LIMIT 5`)
	mock.ExpectQuery(query).WillReturnError(errors.New("failed get products"))

	productRepository := NewProductRepository(db)
	products, err := productRepository.GetBestSellerProducts(ctx)
	assert.NotNil(t, err)
	assert.Nil(t, products)
}

func Test_GetBestSellerProducts_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var eSales = sqlmock.
		NewRows([]string{"ID", "Code", "Name", "Sale"}).
		AddRow(1, "prod-1", "Prod 1", 10000)
	ctx := context.TODO()
	query := regexp.QuoteMeta(`
		SELECT p.ID, p.Code, p.Name, SUM(oi.quantity) as total_sales
			FROM products AS p  JOIN order_items AS oi
			ON p.id = oi.product_id
			GROUP BY oi.product_id
			ORDER BY total_sales DESC
			LIMIT 5`)
	mock.ExpectQuery(query).WillReturnRows(eSales)

	productRepository := NewProductRepository(db)
	aSales, err := productRepository.GetBestSellerProducts(ctx)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eSales, aSales)
}

func Test_GetProductByID_Failed_WhenSelectData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT * FROM products WHERE id = ?")
	mock.ExpectQuery(query).
		WithArgs(int64(1)).
		WillReturnError(errors.New("failed get products"))

	productRepository := NewProductRepository(db)
	products, err := productRepository.GetProductByID(ctx, int64(1))
	assert.NotNil(t, err)
	assert.Nil(t, products)
}

func Test_GetProductByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var eProducts = sqlmock.
		NewRows([]string{"ID", "Code", "Name", "Price", "Stock", "CreatedAt", "UpdatedAt"}).
		AddRow(1, "prod-1", "Prod 1", 10000, 100, time.Now(), time.Now())
	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT * FROM products WHERE id = ?")
	mock.ExpectQuery(query).
		WithArgs(int64(1)).
		WillReturnRows(eProducts)

	productRepository := NewProductRepository(db)
	aProducts, err := productRepository.GetProductByID(ctx, int64(1))
	fmt.Println(aProducts)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eProducts, aProducts)
}

func Test_GetProductByCode_Failed_WhenSelectData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT * FROM products WHERE code = ?")
	mock.ExpectQuery(query).
		WithArgs("prod-1").
		WillReturnError(errors.New("failed get products"))

	productRepository := NewProductRepository(db)
	products, err := productRepository.GetProductByCode(ctx, "prod-1")
	assert.NotNil(t, err)
	assert.Nil(t, products)
}

func Test_GetProductByCode_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var eProducts = sqlmock.
		NewRows([]string{"ID", "Code", "Name", "Price", "Stock", "CreatedAt", "UpdatedAt"}).
		AddRow(1, "prod-1", "Prod 1", 10000, 100, time.Now(), time.Now())
	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT * FROM products WHERE code = ?")
	mock.ExpectQuery(query).
		WithArgs("prod-1").
		WillReturnRows(eProducts)

	productRepository := NewProductRepository(db)
	aProducts, err := productRepository.GetProductByCode(ctx, "prod-1")
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eProducts, aProducts)
}

func Test_CreateProduct_Failed_WhenInsertingProduct(t *testing.T) {
	eProduct := &entity.Product{
		ID:    1,
		Code:  "prod-1",
		Name:  "Prod 1",
		Price: 10000,
		Stock: 100,
	}
	param := entity.CreateProductParam{
		Code:  eProduct.Code,
		Name:  eProduct.Name,
		Price: eProduct.Price,
		Stock: eProduct.Stock,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	queryCreate := regexp.QuoteMeta("INSERT INTO products(code, name, stock, price) VALUES(?, ?, ?, ?)")
	mock.ExpectExec(queryCreate).
		WithArgs(param.Code, param.Name, param.Stock, param.Price).
		WillReturnError(errors.New("failed create product"))

	productRepository := NewProductRepository(db)
	aProducts, err := productRepository.Create(ctx, param)
	assert.NotNil(t, err)
	assert.Nil(t, aProducts)
}

func Test_CreateProduct_Failed_WhenFetchingTheProduct(t *testing.T) {
	eProduct := &entity.Product{
		ID:    1,
		Code:  "prod-1",
		Name:  "Prod 1",
		Price: 10000,
		Stock: 100,
	}
	param := entity.CreateProductParam{
		Code:  eProduct.Code,
		Name:  eProduct.Name,
		Price: eProduct.Price,
		Stock: eProduct.Stock,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	queryCreate := regexp.QuoteMeta("INSERT INTO products(code, name, stock, price) VALUES(?, ?, ?, ?)")
	queryGet := regexp.QuoteMeta("SELECT * FROM products WHERE id = ?")
	mock.ExpectExec(queryCreate).
		WithArgs(param.Code, param.Name, param.Stock, param.Price).
		WillReturnError(errors.New("failed create product"))
	mock.ExpectQuery(queryGet).
		WithArgs(eProduct.ID).
		WillReturnError(errors.New("failed get the product"))

	productRepository := NewProductRepository(db)
	aProducts, err := productRepository.Create(ctx, param)
	assert.NotNil(t, err)
	assert.Nil(t, aProducts)
}

func Test_CreateProduct_Success(t *testing.T) {
	eProduct := &entity.Product{
		ID:        1,
		Code:      "prod-1",
		Name:      "Prod 1",
		Price:     10000,
		Stock:     100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	param := entity.CreateProductParam{
		Code:  eProduct.Code,
		Name:  eProduct.Name,
		Price: eProduct.Price,
		Stock: eProduct.Stock,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	var resProduct = sqlmock.
		NewRows([]string{"ID", "Code", "Name", "Price", "Stock", "CreatedAt", "UpdatedAt"}).
		AddRow(eProduct.ID, eProduct.Code, eProduct.Name, eProduct.Price, eProduct.Stock, eProduct.CreatedAt, eProduct.UpdatedAt)
	queryCreate := regexp.QuoteMeta("INSERT INTO products(code, name, stock, price) VALUES(?, ?, ?, ?)")
	queryGet := regexp.QuoteMeta("SELECT * FROM products WHERE id = ?")
	mock.ExpectExec(queryCreate).
		WithArgs(param.Code, param.Name, param.Stock, param.Price).
		WillReturnResult(sqlmock.NewResult(eProduct.ID, 1))
	mock.ExpectQuery(queryGet).
		WithArgs(eProduct.ID).
		WillReturnRows(resProduct)

	productRepository := NewProductRepository(db)
	aProducts, err := productRepository.Create(ctx, param)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eProduct, aProducts)
}

func Test_UpdateProductByID_Failed(t *testing.T) {
	eProduct := &entity.Product{
		ID:    1,
		Code:  "prod-1-new",
		Name:  "Prod 1 New",
		Price: 12000,
		Stock: 100,
	}
	param := entity.UpdateProductParam{
		Code:  eProduct.Code,
		Name:  eProduct.Name,
		Price: eProduct.Price,
		Stock: eProduct.Stock,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	queryUpdate := regexp.QuoteMeta("UPDATE products SET code = ?, name = ?, stock = ?, price = ? WHERE id = ?")
	mock.ExpectExec(queryUpdate).
		WithArgs(param.Code, param.Name, param.Stock, param.Price).
		WillReturnError(errors.New("failed create product"))

	productRepository := NewProductRepository(db)
	isUpdated, err := productRepository.UpdateByID(ctx, eProduct.ID, param)
	assert.NotNil(t, err)
	assert.False(t, isUpdated)
}

func Test_UpdateProductByID_Success(t *testing.T) {
	eProduct := &entity.Product{
		ID:    1,
		Code:  "prod-1-new",
		Name:  "Prod 1 New",
		Price: 12000,
		Stock: 100,
	}
	param := entity.UpdateProductParam{
		Code:  eProduct.Code,
		Name:  eProduct.Name,
		Price: eProduct.Price,
		Stock: eProduct.Stock,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	queryUpdate := regexp.QuoteMeta("UPDATE products SET code = ?, name = ?, stock = ?, price = ? WHERE id = ?")
	mock.ExpectExec(queryUpdate).
		WithArgs(param.Code, param.Name, param.Stock, param.Price, eProduct.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	productRepository := NewProductRepository(db)
	isUpdated, err := productRepository.UpdateByID(ctx, eProduct.ID, param)
	assert.Nil(t, err)
	assert.True(t, isUpdated)
}

func Test_DecrementProductByIDs_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	usageMap := map[int64]int{1: 2}
	queryUpdate := regexp.QuoteMeta("UPDATE products SET stock = IF(id=?, stock-?, stock) WHERE id IN (?)")
	mock.ExpectExec(queryUpdate).
		WithArgs(1, 2, 1).
		WillReturnError(errors.New("failed decrement products"))

	productRepository := NewProductRepository(db)
	if err := productRepository.DecrementProductByIDs(ctx, usageMap); err != nil {
		assert.NotNil(t, err)
	}
	assert.Nil(t, err)
}

func Test_DecrementProductByIDs_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	usageMap := map[int64]int{1: 2}
	queryUpdate := regexp.QuoteMeta("UPDATE products SET stock = IF(id=?, stock-?, stock) WHERE id IN (?)")
	mock.ExpectExec(queryUpdate).
		WithArgs(1, 2, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	productRepository := NewProductRepository(db)
	if err := productRepository.DecrementProductByIDs(ctx, usageMap); err != nil {
		assert.NotNil(t, err)
	}
	assert.Nil(t, err)
}

func Test_DeleteProductByID_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	productID := int64(1)
	queryUpdate := regexp.QuoteMeta("DELETE FROM products WHERE id = ?")
	mock.ExpectExec(queryUpdate).
		WithArgs(productID).
		WillReturnError(errors.New("failed create product"))

	productRepository := NewProductRepository(db)
	isUpdated, err := productRepository.DeleteByID(ctx, productID)
	assert.NotNil(t, err)
	assert.False(t, isUpdated)
}

func Test_DeleteProductByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	productID := int64(1)
	queryUpdate := regexp.QuoteMeta("DELETE FROM products WHERE id = ?")
	mock.ExpectExec(queryUpdate).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	productRepository := NewProductRepository(db)
	isUpdated, err := productRepository.DeleteByID(ctx, productID)
	assert.Nil(t, err)
	assert.True(t, isUpdated)
}
