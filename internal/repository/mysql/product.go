package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ardafirdausr/go-pos/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(DB *sql.DB) *ProductRepository {
	return &ProductRepository{DB: DB}
}

func (pr ProductRepository) GetAllProducts() ([]*entity.Product, error) {
	rows, err := pr.DB.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	products := []*entity.Product{}
	for rows.Next() {
		var product entity.Product
		var err = rows.Scan(
			&product.ID,
			&product.Code,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		products = append(products, &product)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return products, nil
}

func (pr ProductRepository) GetBestSellerProducts() ([]*entity.ProductSale, error) {
	rows, err := pr.DB.Query(`
		SELECT p.ID, p.Code, p.Name, SUM(oi.quantity) as total_sales
			FROM products AS p  JOIN order_items AS oi
			ON p.id = oi.product_id
			GROUP BY oi.product_id
			ORDER BY total_sales DESC
			LIMIT 5`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	productSales := []*entity.ProductSale{}
	for rows.Next() {
		var productSale entity.ProductSale
		var err = rows.Scan(
			&productSale.ID,
			&productSale.Code,
			&productSale.Name,
			&productSale.Sale)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		productSales = append(productSales, &productSale)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return productSales, nil
}

func (pr ProductRepository) GetProductByCode(code string) (*entity.Product, error) {
	row := pr.DB.QueryRow("SELECT * FROM products WHERE code = ?", code)

	var product entity.Product
	var err = row.Scan(
		&product.ID,
		&product.Code,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		log.Println(err.Error())
		err = entity.ErrNotFound{
			Message: "Product not found",
			Err:     err,
		}
		return nil, err
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &product, nil
}

func (pr ProductRepository) GetProductByID(ID int64) (*entity.Product, error) {
	row := pr.DB.QueryRow("SELECT * FROM products WHERE id = ?", ID)

	var product entity.Product
	var err = row.Scan(
		&product.ID,
		&product.Code,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		log.Println(err.Error())
		err = entity.ErrNotFound{
			Message: "Product not found",
			Err:     err,
		}
		return nil, err
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &product, nil
}

func (pr ProductRepository) GetProductsByIDs(IDs ...int64) ([]*entity.Product, error) {
	if len(IDs) < 1 {
		err := errors.New("ID is required for getting product")
		return nil, err
	}

	IDsString := []string{}
	for _, ID := range IDs {
		IDString := strconv.FormatInt(ID, 10)
		IDsString = append(IDsString, IDString)
	}

	conditionParam := strings.Join(IDsString, ", ")
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s)", conditionParam)
	rows, err := pr.DB.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	products := []*entity.Product{}
	for rows.Next() {
		var product entity.Product
		var err = rows.Scan(
			&product.ID,
			&product.Code,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		products = append(products, &product)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return products, nil
}

func (pr ProductRepository) Create(param entity.CreateProductParam) (*entity.Product, error) {
	res, err := pr.DB.Exec(
		"INSERT INTO products(code, name, stock, price) VALUES(?, ?, ?, ?)",
		param.Code, param.Name, param.Stock, param.Price)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	row := pr.DB.QueryRow("SELECT * FROM products WHERE id = ?", ID)
	err = row.Err()
	if err != nil {
		log.Println(err.Error())
		err = entity.ErrNotFound{Message: "Product not found", Err: err}
		return nil, err
	}

	var product entity.Product
	err = row.Scan(
		&product.ID,
		&product.Code,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pr ProductRepository) UpdateByID(ID int64, param entity.UpdateProductParam) (bool, error) {
	_, err := pr.DB.Exec(
		"UPDATE products SET code = ?, name = ?, stock = ?, price = ? WHERE id = ?",
		param.Code, param.Name, param.Stock, param.Price, ID)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}

func (pr ProductRepository) DeleteByID(ID int64) (bool, error) {
	_, err := pr.DB.Exec("DELETE FROM products WHERE id = ?", ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
