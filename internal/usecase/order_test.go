package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/ardafirdausr/kaseer/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllOrders_Failed(t *testing.T) {
	ctx := context.TODO()
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetAllOrders", ctx).Return(nil, errors.New("failed get orders"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.GetAllOrders(ctx)
	assert.NotNil(t, err)
	assert.Nil(t, aOrders)
}

func Test_GetAllOrders_Success(t *testing.T) {
	eOrders := []*entity.Order{
		{
			ID:    1,
			Total: 40000,
		}, {
			ID:    1,
			Total: 40000,
		},
	}
	ctx := context.TODO()
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetAllOrders", ctx).Return(eOrders, nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.GetAllOrders(ctx)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eOrders, aOrders)
}

func Test_GetOrderItems_Failed(t *testing.T) {
	ctx := context.TODO()
	var orderID int64 = 1
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetOrderItemsByID", ctx, orderID).Return(nil, errors.New("failed get order items"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.GetOrderItems(ctx, orderID)
	assert.NotNil(t, err)
	assert.Nil(t, aOrders)
}

func Test_GetOrderItems_Success(t *testing.T) {
	ctx := context.TODO()
	var orderID int64 = 1
	eOrderItems := make([]*entity.OrderItem, 0)
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetOrderItemsByID", ctx, orderID).Return(eOrderItems, nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrderItems, err := orderUsecase.GetOrderItems(ctx, orderID)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eOrderItems, aOrderItems)
}

func Test_GetAnnualIncome_Failed(t *testing.T) {
	ctx := context.TODO()
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetAnnualIncome", ctx).Return(nil, errors.New("failed get anual income"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetAnnualIncome(ctx)
	assert.NotNil(t, err)
	assert.Nil(t, aRes)
}

func Test_GetAnnualIncome_Success(t *testing.T) {
	ctx := context.TODO()
	eRes := make([]*entity.AnnualIncome, 0)
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetAnnualIncome", ctx).Return(eRes, nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetAnnualIncome(ctx)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(eRes, aRes)
}

func Test_GetDailyOrderCount_Failed(t *testing.T) {
	ctx := context.TODO()
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetDailyOrderCount", ctx).Return(0, errors.New("failed get daily order count"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetDailyOrderCount(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 0, aRes)
}

func Test_GetDailyOrderCount_Success(t *testing.T) {
	ctx := context.TODO()
	eRes := 10
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetDailyOrderCount", ctx).Return(eRes, nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetDailyOrderCount(ctx)
	assert.Nil(t, err)
	assert.Equal(t, eRes, aRes)
}

func Test_GetTotalOrderCount_Failed(t *testing.T) {
	ctx := context.TODO()
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetTotalOrderCount", ctx).Return(0, errors.New("failed get total order count"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetTotalOrderCount(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 0, aRes)
}

func Test_GetTotalOrderCount_Success(t *testing.T) {
	ctx := context.TODO()
	eRes := 10
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetTotalOrderCount", ctx).Return(eRes, nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetTotalOrderCount(ctx)
	assert.Nil(t, err)
	assert.Equal(t, eRes, aRes)
}

func Test_GetLastDayIncome_Failed(t *testing.T) {
	ctx := context.TODO()
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetLastDayIncome", ctx).Return(0, errors.New("failed last daily income"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetLastDayIncome(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 0, aRes)
}

func Test_GetLastDayIncome_Success(t *testing.T) {
	ctx := context.TODO()
	eRes := 10
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetLastDayIncome", ctx).Return(eRes, nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetLastDayIncome(ctx)
	assert.Nil(t, err)
	assert.Equal(t, eRes, aRes)
}

func Test_GetLastMonthIncome_Failed(t *testing.T) {
	ctx := context.TODO()
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetLastMonthIncome", ctx).Return(0, errors.New("failed last month income"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetLastMonthIncome(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, 0, aRes)
}

func Test_GetLastMonthIncome_Success(t *testing.T) {
	ctx := context.TODO()
	eRes := 10
	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("GetLastMonthIncome", ctx).Return(eRes, nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aRes, err := orderUsecase.GetLastMonthIncome(ctx)
	assert.Nil(t, err)
	assert.Equal(t, eRes, aRes)
}

func Test_Create_Failed_WhenCannotGetProductsByID(t *testing.T) {
	ctx := context.TODO()
	var createOrderParam = entity.CreateOrderParam{
		Total: 40000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  10000,
				OrderId:   0,
			}, {
				ProductID: 2,
				Quantity:  3,
				Subtotal:  30000,
				OrderId:   0,
			},
		},
	}

	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductsByIDs", ctx, createOrderParam.Items[0].ProductID, createOrderParam.Items[1].ProductID).Return(nil, errors.New("failed get order items"))
	mockOrderRepo := new(mocks.OrderRepository)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.Create(ctx, createOrderParam)
	assert.NotNil(t, err)
	assert.Nil(t, aOrders)
}

func Test_Create_Failed_WhenQuantityInsufficient(t *testing.T) {
	ctx := context.TODO()
	var createOrderParam = entity.CreateOrderParam{
		Total: 40000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  20000,
				Subtotal:  10000,
				OrderId:   0,
			}, {
				ProductID: 2,
				Quantity:  3,
				Subtotal:  30000,
				OrderId:   0,
			},
		},
	}

	mockUnitOfWork := new(mocks.UnitOfWork)
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductsByIDs", ctx, createOrderParam.Items[0].ProductID, createOrderParam.Items[1].ProductID).Return(products, nil)
	mockOrderRepo := new(mocks.OrderRepository)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.Create(ctx, createOrderParam)
	assert.NotNil(t, err)
	assert.IsType(t, err, entity.ErrValidation{})
	assert.Nil(t, aOrders)
}

func Test_Create_Failed_WhenBeginTransaction(t *testing.T) {
	ctx := context.TODO()
	var createOrderParam = entity.CreateOrderParam{
		Total: 40000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  10000,
				OrderId:   0,
			}, {
				ProductID: 2,
				Quantity:  3,
				Subtotal:  30000,
				OrderId:   0,
			},
		},
	}

	mockUnitOfWork := new(mocks.UnitOfWork)
	mockUnitOfWork.On("Begin", ctx).Return(nil, errors.New("expired context"))
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductsByIDs", ctx, createOrderParam.Items[0].ProductID, createOrderParam.Items[1].ProductID).Return(products, nil)
	mockOrderRepo := new(mocks.OrderRepository)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.Create(ctx, createOrderParam)
	assert.NotNil(t, err)
	assert.Nil(t, aOrders)
}

func Test_Create_Failed_WhenCreatingOrder(t *testing.T) {
	ctx := context.TODO()
	var createOrderParam = entity.CreateOrderParam{
		Total: 40000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  10000,
				OrderId:   0,
			}, {
				ProductID: 2,
				Quantity:  3,
				Subtotal:  30000,
				OrderId:   0,
			},
		},
	}

	mockUnitOfWork := new(mocks.UnitOfWork)
	mockUnitOfWork.On("Begin", ctx).Return(ctx, nil)
	mockUnitOfWork.On("Rollback", ctx).Return(nil)
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductsByIDs", ctx, createOrderParam.Items[0].ProductID, createOrderParam.Items[1].ProductID).Return(products, nil)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("Create", ctx, createOrderParam).Return(nil, errors.New("failed creating order"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.Create(ctx, createOrderParam)
	assert.NotNil(t, err)
	assert.Nil(t, aOrders)
}

func Test_Create_Failed_WhenCreatingOrderItems(t *testing.T) {
	ctx := context.TODO()
	var createOrderParam = entity.CreateOrderParam{
		Total: 40000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  10000,
				OrderId:   0,
			}, {
				ProductID: 2,
				Quantity:  3,
				Subtotal:  30000,
				OrderId:   0,
			},
		},
	}
	var eOrder = &entity.Order{
		ID:    1,
		Total: createOrderParam.Total,
	}

	mockUnitOfWork := new(mocks.UnitOfWork)
	mockUnitOfWork.On("Begin", ctx).Return(ctx, nil)
	mockUnitOfWork.On("Rollback", ctx).Return(nil)
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductsByIDs", ctx, createOrderParam.Items[0].ProductID, createOrderParam.Items[1].ProductID).Return(products, nil)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("Create", ctx, createOrderParam).Return(eOrder, nil)
	mockOrderRepo.On("CreateOrderItems", ctx, eOrder.ID, createOrderParam.Items).Return(errors.New("failed create order items"))

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.Create(ctx, createOrderParam)
	assert.NotNil(t, err)
	assert.Nil(t, aOrders)
}

func Test_Create_Failed_WhenDecreasingProductQuantity(t *testing.T) {
	ctx := context.TODO()
	var createOrderParam = entity.CreateOrderParam{
		Total: 40000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  10000,
				OrderId:   0,
			}, {
				ProductID: 2,
				Quantity:  3,
				Subtotal:  30000,
				OrderId:   0,
			},
		},
	}
	var productSale = map[int64]int{1: 2, 2: 3}
	var eOrder = &entity.Order{
		ID:    1,
		Total: createOrderParam.Total,
	}

	mockUnitOfWork := new(mocks.UnitOfWork)
	mockUnitOfWork.On("Begin", ctx).Return(ctx, nil)
	mockUnitOfWork.On("Rollback", ctx).Return(nil)
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductsByIDs", ctx, createOrderParam.Items[0].ProductID, createOrderParam.Items[1].ProductID).Return(products, nil)
	mockProductRepo.On("DecrementProductByIDs", ctx, productSale).Return(errors.New("failed to decrease product quantity"))
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("Create", ctx, createOrderParam).Return(eOrder, nil)
	mockOrderRepo.On("CreateOrderItems", ctx, eOrder.ID, createOrderParam.Items).Return(nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.Create(ctx, createOrderParam)
	assert.NotNil(t, err)
	assert.Nil(t, aOrders)
}

func Test_Create_Failed_WhenCommitingTransaction(t *testing.T) {
	ctx := context.TODO()
	var createOrderParam = entity.CreateOrderParam{
		Total: 40000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  10000,
				OrderId:   0,
			}, {
				ProductID: 2,
				Quantity:  3,
				Subtotal:  30000,
				OrderId:   0,
			},
		},
	}
	var productSale = map[int64]int{1: 2, 2: 3}
	var eOrder = &entity.Order{
		ID:    1,
		Total: createOrderParam.Total,
	}

	mockUnitOfWork := new(mocks.UnitOfWork)
	mockUnitOfWork.On("Begin", ctx).Return(ctx, nil)
	mockUnitOfWork.On("Commit", ctx).Return(errors.New("failed to commit transcation"))
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductsByIDs", ctx, createOrderParam.Items[0].ProductID, createOrderParam.Items[1].ProductID).Return(products, nil)
	mockProductRepo.On("DecrementProductByIDs", ctx, productSale).Return(nil)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("Create", ctx, createOrderParam).Return(eOrder, nil)
	mockOrderRepo.On("CreateOrderItems", ctx, eOrder.ID, createOrderParam.Items).Return(nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrders, err := orderUsecase.Create(ctx, createOrderParam)
	assert.NotNil(t, err)
	assert.Nil(t, aOrders)
}

func Test_Create_Success(t *testing.T) {
	ctx := context.TODO()
	var createOrderParam = entity.CreateOrderParam{
		Total: 40000,
		Items: []*entity.CreateOrderItemParam{
			{
				ProductID: 1,
				Quantity:  2,
				Subtotal:  10000,
				OrderId:   0,
			}, {
				ProductID: 2,
				Quantity:  3,
				Subtotal:  30000,
				OrderId:   0,
			},
		},
	}
	var productSale = map[int64]int{1: 2, 2: 3}
	var eOrder = &entity.Order{
		ID:    1,
		Total: createOrderParam.Total,
	}

	mockUnitOfWork := new(mocks.UnitOfWork)
	mockUnitOfWork.On("Begin", ctx).Return(ctx, nil)
	mockUnitOfWork.On("Commit", ctx).Return(nil)
	mockProductRepo := new(mocks.ProductRepository)
	mockProductRepo.On("GetProductsByIDs", ctx, createOrderParam.Items[0].ProductID, createOrderParam.Items[1].ProductID).Return(products, nil)
	mockProductRepo.On("DecrementProductByIDs", ctx, productSale).Return(nil)
	mockOrderRepo := new(mocks.OrderRepository)
	mockOrderRepo.On("Create", ctx, createOrderParam).Return(eOrder, nil)
	mockOrderRepo.On("CreateOrderItems", ctx, eOrder.ID, createOrderParam.Items).Return(nil)

	orderUsecase := NewOrderUsecase(mockOrderRepo, mockProductRepo, mockUnitOfWork)
	aOrder, err := orderUsecase.Create(ctx, createOrderParam)
	assert.Nil(t, err)
	assert.ObjectsAreEqual(eOrder, aOrder)
}
