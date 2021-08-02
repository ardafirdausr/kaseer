package mysql

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllOrders_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT * from orders ORDER BY created_at DESC")
	mock.ExpectQuery(query).WillReturnError(errors.New("failed get orders"))

	OrderRepository := NewOrderRepository(db)
	orders, err := OrderRepository.GetAllOrders(ctx)
	assert.NotNil(t, err)
	assert.Nil(t, orders)
}

func Test_GetAllOrders_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var eOrders = sqlmock.
		NewRows([]string{"ID", "Total", "CreatedAt"}).
		AddRow(1, 20000, time.Now())
	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT * from orders ORDER BY created_at DESC")
	mock.ExpectQuery(query).WillReturnRows(eOrders)

	OrderRepository := NewOrderRepository(db)
	aOrders, err := OrderRepository.GetAllOrders(ctx)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eOrders, aOrders)
}

func Test_GetAnnualIncome_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta(`
		SELECT YEAR(created_at) as year, MONTHNAME(created_at) as mount, SUM(total) as income
			FROM orders
			WHERE MONTH(created_at) -12 AND MONTH(created_at)
			GROUP BY YEAR(created_at), MONTHNAME(created_at), MONTH(created_at)
			ORDER BY YEAR(created_at) ASC, MONTH(created_at) ASC`)
	mock.ExpectQuery(query).WillReturnError(errors.New("failed get anual income"))

	OrderRepository := NewOrderRepository(db)
	income, err := OrderRepository.GetAnnualIncome(ctx)
	assert.NotNil(t, err)
	assert.Nil(t, income)
}

func Test_GetAnnualIncome_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var eIncome = sqlmock.
		NewRows([]string{"Year", "Month", "Income"}).
		AddRow(2021, "January", 1400000).
		AddRow(2021, "February", 2000000)
	ctx := context.TODO()
	query := regexp.QuoteMeta(`
		SELECT YEAR(created_at) as year, MONTHNAME(created_at) as mount, SUM(total) as income
			FROM orders
			WHERE MONTH(created_at) -12 AND MONTH(created_at)
			GROUP BY YEAR(created_at), MONTHNAME(created_at), MONTH(created_at)
			ORDER BY YEAR(created_at) ASC, MONTH(created_at) ASC`)
	mock.ExpectQuery(query).WillReturnRows(eIncome)

	OrderRepository := NewOrderRepository(db)
	aIncome, err := OrderRepository.GetAnnualIncome(ctx)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eIncome, aIncome)
}

func Test_GetDailyOrderCount_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT COUNT(*) FROM orders WHERE DAY(created_At) = DAY(CURRENT_TIMESTAMP())")
	mock.ExpectQuery(query).WillReturnError(errors.New("failed get daily order"))

	OrderRepository := NewOrderRepository(db)
	res, err := OrderRepository.GetDailyOrderCount(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 0, res)
}

func Test_GetDailyOrderCount_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var order = sqlmock.NewRows([]string{""}).AddRow(0)
	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT COUNT(*) FROM orders WHERE DAY(created_At) = DAY(CURRENT_TIMESTAMP())")
	mock.ExpectQuery(query).WillReturnRows(order)

	OrderRepository := NewOrderRepository(db)
	res, err := OrderRepository.GetDailyOrderCount(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, res)
}

func Test_GetTotalOrderCount_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT COUNT(*) FROM orders")
	mock.ExpectQuery(query).WillReturnError(errors.New("failed get total orders"))

	OrderRepository := NewOrderRepository(db)
	res, err := OrderRepository.GetTotalOrderCount(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 0, res)
}

func Test_GetTotalOrderCount_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var order = sqlmock.NewRows([]string{""}).AddRow(0)
	ctx := context.TODO()
	query := regexp.QuoteMeta("SELECT COUNT(*) FROM orders")
	mock.ExpectQuery(query).WillReturnRows(order)

	OrderRepository := NewOrderRepository(db)
	res, err := OrderRepository.GetTotalOrderCount(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, res)
}

func Test_GetLastDayIncome_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta(`
		SELECT SUM(total)
			FROM orders
			WHERE DAY(created_At) = DAY(CURRENT_TIMESTAMP())
			GROUP BY DAY(created_At)`)
	mock.ExpectQuery(query).WillReturnError(errors.New("failed get last daily income orders"))

	OrderRepository := NewOrderRepository(db)
	res, err := OrderRepository.GetLastDayIncome(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 0, res)
}

func Test_GetLastDayIncome_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var order = sqlmock.NewRows([]string{""}).AddRow(0)
	ctx := context.TODO()
	query := regexp.QuoteMeta(`
		SELECT SUM(total)
			FROM orders
			WHERE DAY(created_At) = DAY(CURRENT_TIMESTAMP())
			GROUP BY DAY(created_At)`)
	mock.ExpectQuery(query).WillReturnRows(order)

	OrderRepository := NewOrderRepository(db)
	res, err := OrderRepository.GetLastDayIncome(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, res)
}

func Test_GetLastMonthIncome_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	query := regexp.QuoteMeta(`
		SELECT SUM(total)
			FROM orders
			WHERE MONTH(created_At) = MONTH(CURRENT_TIMESTAMP())
			GROUP BY MONTH(created_At)`)
	mock.ExpectQuery(query).WillReturnError(errors.New("failed get last month income orders"))

	OrderRepository := NewOrderRepository(db)
	res, err := OrderRepository.GetLastMonthIncome(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 0, res)
}

func Test_GetLastMonthIncome_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var order = sqlmock.NewRows([]string{""}).AddRow(0)
	ctx := context.TODO()
	query := regexp.QuoteMeta(`
		SELECT SUM(total)
			FROM orders
			WHERE MONTH(created_At) = MONTH(CURRENT_TIMESTAMP())
			GROUP BY MONTH(created_At)`)
	mock.ExpectQuery(query).WillReturnRows(order)

	OrderRepository := NewOrderRepository(db)
	res, err := OrderRepository.GetLastMonthIncome(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, res)
}

func Test_GetOrderItemsByID_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	orderID := int64(1)
	query := regexp.QuoteMeta(`
		SELECT p.id as id, oi.order_id, p.id as product_id, p.code, p.name, p.price, oi.quantity, oi.subtotal, oi.created_at
				FROM order_items AS oi
				LEFT JOIN products AS p ON oi.product_id = p.id
				WHERE oi.order_id = ?`)
	mock.ExpectQuery(query).
		WithArgs(orderID).
		WillReturnError(errors.New("failed get orders"))

	OrderRepository := NewOrderRepository(db)
	orders, err := OrderRepository.GetOrderItemsByID(ctx, orderID)
	assert.NotNil(t, err)
	assert.Nil(t, orders)
}

func Test_GetOrderItemsByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var eOrderItems = sqlmock.
		NewRows([]string{"ID", "OrderID", "ProductID", "ProductCode", "ProductName", "ProductPrice", "Quantity", "Subtotal", "CreatedAt"}).
		AddRow(1, 1, 1, "prod-1", "Prod 1", 10000, 2, 10000, time.Now()).
		AddRow(2, 1, 2, "prod-2", "Prod 2", 15000, 2, 30000, time.Now())
	ctx := context.TODO()
	orderID := int64(1)
	query := regexp.QuoteMeta(`
		SELECT p.id as id, oi.order_id, p.id as product_id, p.code, p.name, p.price, oi.quantity, oi.subtotal, oi.created_at
				FROM order_items AS oi
				LEFT JOIN products AS p ON oi.product_id = p.id
				WHERE oi.order_id = ?`)
	mock.ExpectQuery(query).
		WithArgs(orderID).
		WillReturnRows(eOrderItems)

	OrderRepository := NewOrderRepository(db)
	aOrderItems, err := OrderRepository.GetOrderItemsByID(ctx, orderID)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eOrderItems, aOrderItems)
}

func Test_CreateOrder_Failed(t *testing.T) {
	param := entity.CreateOrderParam{
		Total: 50000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  20000,
				OrderId:   1,
			}, {
				ProductID: 2,
				Quantity:  2,
				Subtotal:  30000,
				OrderId:   1,
			},
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	queryCreate := regexp.QuoteMeta("INSERT INTO orders(total) VALUES(?)")
	mock.ExpectExec(queryCreate).
		WithArgs(param.Total).
		WillReturnError(errors.New("failed create order"))

	OrderRepository := NewOrderRepository(db)
	aProducts, err := OrderRepository.Create(ctx, param)
	assert.NotNil(t, err)
	assert.Nil(t, aProducts)
}

func Test_CreateOrder_Success(t *testing.T) {
	eOrder := &entity.Order{
		ID:    1,
		Total: 50000,
	}
	param := entity.CreateOrderParam{
		Total: 50000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  20000,
				OrderId:   1,
			}, {
				ProductID: 2,
				Quantity:  2,
				Subtotal:  30000,
				OrderId:   1,
			},
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	queryCreate := regexp.QuoteMeta("INSERT INTO orders(total) VALUES(?)")
	mock.ExpectExec(queryCreate).
		WithArgs(param.Total).
		WillReturnResult(sqlmock.NewResult(1, 1))

	OrderRepository := NewOrderRepository(db)
	aOrder, err := OrderRepository.Create(ctx, param)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eOrder, aOrder)
}

func Test_CreateOrderItems_Failed(t *testing.T) {
	orderID := int64(1)
	param := []*entity.CreateOrderItemParam{
		{
			ProductID: 1,
			Quantity:  2,
			Subtotal:  20000,
			OrderId:   1,
		}, {
			ProductID: 2,
			Quantity:  2,
			Subtotal:  30000,
			OrderId:   1,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	queryCreate := regexp.QuoteMeta("INSERT INTO order_items(order_id, product_id, quantity, subtotal) VALUES (?, ?, ?, ?), (?, ?, ?, ?)")
	mock.ExpectExec(queryCreate).
		WithArgs(
			param[0].OrderId, param[0].ProductID, param[0].Quantity, param[0].Subtotal,
			param[1].OrderId, param[1].ProductID, param[1].Quantity, param[1].Subtotal,
		).
		WillReturnError(errors.New("failed create order items"))

	OrderRepository := NewOrderRepository(db)
	if err := OrderRepository.CreateOrderItems(ctx, orderID, param); err != nil {
		assert.NotNil(t, err)
		assert.Equal(t, "failed create order items", err.Error())
	}
	assert.Nil(t, err)
}

func Test_CreateOrderItems_Success(t *testing.T) {
	orderID := int64(1)
	param := []*entity.CreateOrderItemParam{
		{
			ProductID: 1,
			Quantity:  2,
			Subtotal:  20000,
			OrderId:   1,
		}, {
			ProductID: 2,
			Quantity:  2,
			Subtotal:  30000,
			OrderId:   1,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.TODO()
	queryCreate := regexp.QuoteMeta("INSERT INTO order_items(order_id, product_id, quantity, subtotal) VALUES (?, ?, ?, ?), (?, ?, ?, ?)")
	mock.ExpectExec(queryCreate).
		WithArgs(
			param[0].OrderId, param[0].ProductID, param[0].Quantity, param[0].Subtotal,
			param[1].OrderId, param[1].ProductID, param[1].Quantity, param[1].Subtotal,
		).
		WillReturnResult(sqlmock.NewResult(2, 2))

	OrderRepository := NewOrderRepository(db)
	if err := OrderRepository.CreateOrderItems(ctx, orderID, param); err != nil {
		assert.NotNil(t, err)
	}
	assert.Nil(t, err)
}
