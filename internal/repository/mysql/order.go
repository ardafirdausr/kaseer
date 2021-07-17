package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ardafirdausr/kaseer/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(DB *sql.DB) *OrderRepository {
	return &OrderRepository{DB: DB}
}

func (or OrderRepository) GetAllOrders() ([]*entity.Order, error) {
	rows, err := or.DB.Query("SELECT * from orders ORDER BY created_at DESC")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	orders := []*entity.Order{}
	for rows.Next() {
		var order entity.Order
		var err = rows.Scan(
			&order.ID,
			&order.Total,
			&order.CreatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		orders = append(orders, &order)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return orders, nil
}

func (or OrderRepository) GetAnnualIncome() ([]*entity.AnnualIncome, error) {
	query := `
		SELECT YEAR(created_at) as year, MONTHNAME(created_at) as mount, SUM(total) as income
			FROM orders
			WHERE MONTH(created_at) -12 AND MONTH(created_at)
			GROUP BY YEAR(created_at), MONTHNAME(created_at), MONTH(created_at)
			ORDER BY YEAR(created_at) ASC, MONTH(created_at) ASC`
	rows, err := or.DB.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	incomes := []*entity.AnnualIncome{}
	for rows.Next() {
		var income entity.AnnualIncome
		var err = rows.Scan(&income.Year, &income.Month, &income.Income)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		incomes = append(incomes, &income)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return incomes, nil
}

func (or OrderRepository) GetDailyOrderCount() (int, error) {
	row := or.DB.QueryRow("SELECT COUNT(*) FROM orders WHERE DAY(created_At) = DAY(CURRENT_TIMESTAMP())")
	if err := row.Err(); err != nil {
		return 0, err
	}

	var val int
	if err := row.Scan(&val); err != nil {
		return 0, err
	}

	return val, nil
}

func (or OrderRepository) GetTotalOrderCount() (int, error) {
	row := or.DB.QueryRow("SELECT COUNT(*) FROM orders")
	if err := row.Err(); err != nil {
		return 0, err
	}

	var val int
	if err := row.Scan(&val); err != nil {
		return 0, err
	}

	return val, nil
}

func (or OrderRepository) GetLastDayIncome() (int, error) {
	query := `
		SELECT SUM(total)
			FROM orders
			WHERE DAY(created_At) = DAY(CURRENT_TIMESTAMP())
			GROUP BY DAY(created_At)`
	row := or.DB.QueryRow(query)
	var val int
	err := row.Scan(&val)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return val, nil
}

func (or OrderRepository) GetLastMonthIncome() (int, error) {
	query := `
		SELECT SUM(total)
			FROM orders
			WHERE MONTH(created_At) = MONTH(CURRENT_TIMESTAMP())
			GROUP BY MONTH(created_At)`
	row := or.DB.QueryRow(query)
	var val int
	err := row.Scan(&val)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return val, nil
}

func (or OrderRepository) GetOrderItemsByID(ID int64) ([]*entity.OrderItem, error) {
	rows, err := or.DB.Query(`
		SELECT p.id as id, oi.order_id, p.id as product_id, p.code, p.name, p.price, oi.quantity, oi.subtotal, oi.created_at
			FROM order_items AS oi
			LEFT JOIN products AS p ON oi.product_id = p.id
			WHERE oi.order_id = ?`,
		ID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	orderItems := []*entity.OrderItem{}
	for rows.Next() {
		var orderItem entity.OrderItem
		err := rows.Scan(
			&orderItem.ID,
			&orderItem.OrderID,
			&orderItem.ProductID,
			&orderItem.ProductCode,
			&orderItem.ProductName,
			&orderItem.ProductPrice,
			&orderItem.Quantity,
			&orderItem.Subtotal,
			&orderItem.CreatedAt,
		)
		if err != nil {
			log.Println(err.Error())
		}

		orderItems = append(orderItems, &orderItem)
	}

	return orderItems, nil
}

func (or OrderRepository) Create(param entity.CreateOrderParam) (*entity.Order, error) {
	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := or.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// create orders
	ors, err := tx.ExecContext(ctx, "INSERT INTO orders(total) VALUES(?)", param.Total)
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, err
	}

	orderId, err := ors.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, err
	}

	// create order items
	createOrderParams := []string{}
	createOrderVals := []interface{}{}
	for _, item := range param.Items {
		createOrderParams = append(createOrderParams, "(?, ?, ?, ?)")
		createOrderVals = append(createOrderVals, orderId, item.ProductId, item.Quantity, item.Subtotal)
		item.OrderId = orderId
	}

	createOrderParamQuery := strings.Join(createOrderParams, ", ")
	query := fmt.Sprintf("INSERT INTO order_items(order_id, product_id, quantity, subtotal) VALUES %s", createOrderParamQuery)
	_, err = tx.ExecContext(ctx, query, createOrderVals...)
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, err
	}

	// Decrement stocks
	decrementStockParams := []string{}
	decrementProductIDs := []string{}
	for _, item := range param.Items {
		decrementProductIDs = append(decrementProductIDs, strconv.FormatInt(item.ProductId, 10))
		param := fmt.Sprintf("stock = IF(id=%d, stock-%d, stock)", item.ProductId, item.Quantity)
		decrementStockParams = append(decrementStockParams, param)
	}

	decrementStockParamQuery := strings.Join(decrementStockParams, ", ")
	decrementProductIDsQuery := strings.Join(decrementProductIDs, ", ")
	query = fmt.Sprintf("UPDATE products SET %s WHERE id IN (%s)", decrementStockParamQuery, decrementProductIDsQuery)
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Println(err.Error())
	}

	return nil, nil
}
