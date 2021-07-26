package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/entity"
)

type OrderUsecase struct {
	orderRepository   internal.OrderRepository
	productRepository internal.ProductRepository
	UnitOfWork        internal.UnitOfWork
}

func NewOrderUsecase(
	orderRepository internal.OrderRepository,
	productRepository internal.ProductRepository,
	UnitOfWork internal.UnitOfWork) *OrderUsecase {
	return &OrderUsecase{orderRepository, productRepository, UnitOfWork}
}

func (ou OrderUsecase) GetAllOrders(ctx context.Context) ([]*entity.Order, error) {
	orders, err := ou.orderRepository.GetAllOrders(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	return orders, err
}

func (ou OrderUsecase) GetOrderItems(ctx context.Context, orderID int64) ([]*entity.OrderItem, error) {
	orderItems, err := ou.orderRepository.GetOrderItemsByID(ctx, orderID)
	if err != nil {
		log.Println(err.Error())
	}

	return orderItems, err
}

func (ou OrderUsecase) GetAnnualIncome(ctx context.Context) ([]*entity.AnnualIncome, error) {
	annualIncomes, err := ou.orderRepository.GetAnnualIncome(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	return annualIncomes, err
}

func (ou OrderUsecase) GetDailyOrderCount(ctx context.Context) (int, error) {
	res, err := ou.orderRepository.GetDailyOrderCount(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	return res, err
}

func (ou OrderUsecase) GetTotalOrderCount(ctx context.Context) (int, error) {
	res, err := ou.orderRepository.GetTotalOrderCount(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	return res, err
}

func (ou OrderUsecase) GetLastDayIncome(ctx context.Context) (int, error) {
	res, err := ou.orderRepository.GetLastDayIncome(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	return res, err
}

func (ou OrderUsecase) GetLastMonthIncome(ctx context.Context) (int, error) {
	res, err := ou.orderRepository.GetLastMonthIncome(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	return res, err
}

func (ou OrderUsecase) Create(ctx context.Context, param entity.CreateOrderParam) (*entity.Order, error) {
	// check available quantity
	orderQuantity := make(map[int64]int)
	productSale := make(map[int64]int)
	productIDs := make([]int64, 0)

	for _, item := range param.Items {
		orderQuantity[item.ProductID] = item.Quantity
		productSale[item.ProductID] = item.Quantity
		productIDs = append(productIDs, item.ProductID)
	}

	products, err := ou.productRepository.GetProductsByIDs(ctx, productIDs...)
	if err != nil {
		return nil, err
	}

	ev := entity.ErrValidation{
		Message: "Insufficient product quantity",
		Errors:  map[string]string{},
	}
	for _, product := range products {
		if product.Stock-orderQuantity[product.ID] < 0 {
			ev.Errors[product.Name] = fmt.Sprintf("%s remaining quantity: %d", product.Name, product.Stock)
		}
	}

	if len(ev.Errors) > 0 {
		return nil, ev
	}

	txContext, err := ou.UnitOfWork.Begin(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	order, err := ou.orderRepository.Create(txContext, param)
	if err != nil {
		log.Println(err.Error())
		ou.UnitOfWork.Rollback(txContext)
		return nil, err
	}

	if err := ou.orderRepository.CreateOrderItems(txContext, order.ID, param.Items); err != nil {
		log.Println(err.Error())
		ou.UnitOfWork.Rollback(txContext)
		return nil, err
	}

	if err := ou.productRepository.DecrementProductByIDs(txContext, productSale); err != nil {
		log.Println(err.Error())
		ou.UnitOfWork.Rollback(txContext)
		return nil, err
	}

	if err := ou.UnitOfWork.Commit(txContext); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return order, nil
}
