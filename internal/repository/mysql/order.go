package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ardafirdausr/kaseer/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(DB *sql.DB) *OrderRepository {
	return &OrderRepository{DB: DB}
}

func (repo OrderRepository) GetAllOrders(ctx context.Context) ([]*entity.Order, error) {
	var rows *sql.Rows
	var err error
	query := "SELECT * from orders ORDER BY created_at DESC"
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

func (repo OrderRepository) GetAnnualIncome(ctx context.Context) ([]*entity.AnnualIncome, error) {
	var rows *sql.Rows
	var err error
	query := `
		SELECT YEAR(created_at) as year, MONTHNAME(created_at) as mount, SUM(total) as income
			FROM orders
			WHERE MONTH(created_at) -12 AND MONTH(created_at)
			GROUP BY YEAR(created_at), MONTHNAME(created_at), MONTH(created_at)
			ORDER BY YEAR(created_at) ASC, MONTH(created_at) ASC`
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

func (repo OrderRepository) GetDailyOrderCount(ctx context.Context) (int, error) {
	var row *sql.Row
	query := "SELECT COUNT(*) FROM orders WHERE DAY(created_At) = DAY(CURRENT_TIMESTAMP())"
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		row = tx.QueryRow(query)
	} else {
		row = repo.DB.QueryRowContext(ctx, query)
	}

	if err := row.Err(); err != nil {
		return 0, err
	}

	var val int
	if err := row.Scan(&val); err != nil {
		return 0, err
	}

	return val, nil
}

func (repo OrderRepository) GetTotalOrderCount(ctx context.Context) (int, error) {
	var row *sql.Row
	query := "SELECT COUNT(*) FROM orders"
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		row = tx.QueryRow(query)
	} else {
		row = repo.DB.QueryRowContext(ctx, query)
	}

	if err := row.Err(); err != nil {
		return 0, err
	}

	var val int
	if err := row.Scan(&val); err != nil {
		return 0, err
	}

	return val, nil
}

func (repo OrderRepository) GetLastDayIncome(ctx context.Context) (int, error) {
	var row *sql.Row
	query := `
		SELECT SUM(total)
			FROM orders
			WHERE DAY(created_At) = DAY(CURRENT_TIMESTAMP())
			GROUP BY DAY(created_At)`
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		row = tx.QueryRow(query)
	} else {
		row = repo.DB.QueryRowContext(ctx, query)
	}

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

func (repo OrderRepository) GetLastMonthIncome(ctx context.Context) (int, error) {
	var row *sql.Row
	query := `
		SELECT SUM(total)
			FROM orders
			WHERE MONTH(created_At) = MONTH(CURRENT_TIMESTAMP())
			GROUP BY MONTH(created_At)`
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		row = tx.QueryRow(query)
	} else {
		row = repo.DB.QueryRowContext(ctx, query)
	}

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

func (repo OrderRepository) GetOrderItemsByID(ctx context.Context, ID int64) ([]*entity.OrderItem, error) {
	var rows *sql.Rows
	var err error
	query := `
		SELECT p.id as id, oi.order_id, p.id as product_id, p.code, p.name, p.price, oi.quantity, oi.subtotal, oi.created_at
				FROM order_items AS oi
				LEFT JOIN products AS p ON oi.product_id = p.id
				WHERE oi.order_id = ?`
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		rows, err = tx.Query(query, ID)
	} else {
		rows, err = repo.DB.QueryContext(ctx, query, ID)
	}

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

func (repo OrderRepository) Create(ctx context.Context, param entity.CreateOrderParam) (*entity.Order, error) {
	query := "INSERT INTO orders(total) VALUES(?)"
	var res sql.Result
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		res, err = tx.Exec(query, param.Total)
	} else {
		res, err = repo.DB.ExecContext(ctx, query, param.Total)
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

	order := &entity.Order{
		ID:        ID,
		Total:     param.Total,
		CreatedAt: time.Now(),
	}
	return order, nil
}

func (repo OrderRepository) CreateOrderItems(ctx context.Context, orderID int64, items []*entity.CreateOrderItemParam) error {
	createOrderParams := []string{}
	createOrderVals := []interface{}{}
	for _, item := range items {
		createOrderParams = append(createOrderParams, "(?, ?, ?, ?)")
		createOrderVals = append(createOrderVals, orderID, item.ProductID, item.Quantity, item.Subtotal)
	}
	createOrderParamQuery := strings.Join(createOrderParams, ", ")

	query := fmt.Sprintf("INSERT INTO order_items(order_id, product_id, quantity, subtotal) VALUES %s", createOrderParamQuery)
	var err error
	if tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx); ok {
		_, err = tx.Exec(query, createOrderVals...)
	} else {
		_, err = repo.DB.ExecContext(ctx, query, createOrderVals...)
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
