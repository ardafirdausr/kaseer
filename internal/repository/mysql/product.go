package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ardafirdausr/kaseer/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(DB *sql.DB) *ProductRepository {
	return &ProductRepository{DB: DB}
}

func (repo ProductRepository) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	var rows *sql.Rows
	var err error
	query := "SELECT * FROM products"
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		rows, err = tx.Query(query)
	} else {
		rows, err = repo.DB.QueryContext(ctx, query)
	}

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

func (repo ProductRepository) GetBestSellerProducts(ctx context.Context) ([]*entity.ProductSale, error) {
	var rows *sql.Rows
	var err error
	query := `
		SELECT p.ID, p.Code, p.Name, SUM(oi.quantity) as total_sales
			FROM products AS p  JOIN order_items AS oi
			ON p.id = oi.product_id
			GROUP BY oi.product_id
			ORDER BY total_sales DESC
			LIMIT 5`
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		rows, err = tx.Query(query)
	} else {
		rows, err = repo.DB.QueryContext(ctx, query)
	}

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

func (repo ProductRepository) GetProductByCode(ctx context.Context, code string) (*entity.Product, error) {
	var row *sql.Row
	query := "SELECT * FROM products WHERE code = ?"
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		row = tx.QueryRow(query, code)
	} else {
		row = repo.DB.QueryRowContext(ctx, query, code)
	}

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

func (repo ProductRepository) GetProductByID(ctx context.Context, ID int64) (*entity.Product, error) {
	var row *sql.Row
	query := "SELECT * FROM products WHERE id = ?"
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		row = tx.QueryRow(query, ID)
	} else {
		row = repo.DB.QueryRowContext(ctx, query, ID)
	}

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

func (repo ProductRepository) GetProductsByIDs(ctx context.Context, IDs ...int64) ([]*entity.Product, error) {
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
	var rows *sql.Rows
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		rows, err = tx.Query(query)
	} else {
		rows, err = repo.DB.QueryContext(ctx, query)
	}

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

func (repo ProductRepository) Create(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error) {
	query := "INSERT INTO products(code, name, stock, price) VALUES(?, ?, ?, ?)"
	var res sql.Result
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		res, err = tx.Exec(query, param.Code, param.Name, param.Stock, param.Price)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, param.Code, param.Name, param.Stock, param.Price)
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	row := repo.DB.QueryRowContext(ctx, "SELECT * FROM products WHERE id = ?", ID)
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

func (repo ProductRepository) UpdateByID(ctx context.Context, ID int64, param entity.UpdateProductParam) (bool, error) {
	query := "UPDATE products SET code = ?, name = ?, stock = ?, price = ? WHERE id = ?"
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		_, err = tx.Exec(query, param.Code, param.Name, param.Stock, param.Price, ID)
	} else {
		_, err = repo.DB.ExecContext(ctx, query, param.Code, param.Name, param.Stock, param.Price, ID)
	}

	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	return true, nil
}

func (repo ProductRepository) DecrementProductByIDs(ctx context.Context, IDDecrementMap map[int64]int) error {
	decrementStockParams := []string{}
	decrementProductIDs := []string{}
	for id, quantity := range IDDecrementMap {
		decrementProductIDs = append(decrementProductIDs, strconv.FormatInt(id, 10))
		param := fmt.Sprintf("stock = IF(id=%d, stock-%d, stock)", id, quantity)
		decrementStockParams = append(decrementStockParams, param)
	}

	decrementStockParamQuery := strings.Join(decrementStockParams, ", ")
	decrementProductIDsQuery := strings.Join(decrementProductIDs, ", ")
	query := fmt.Sprintf("UPDATE products SET %s WHERE id IN (%s)", decrementStockParamQuery, decrementProductIDsQuery)
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		_, err = tx.Exec(query)
	} else {
		_, err = repo.DB.ExecContext(ctx, query)
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (repo ProductRepository) DeleteByID(ctx context.Context, ID int64) (bool, error) {
	query := "DELETE FROM products WHERE id = ?"
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		_, err = tx.Exec(query, ID)
	} else {
		_, err = repo.DB.ExecContext(ctx, query, ID)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
