package usecase

import (
	"fmt"
	"log"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/entity"
)

type OrderUsecase struct {
	orderRepository   internal.OrderRepository
	productRepository internal.ProductRepository
}

func NewOrderUsecase(orderRepository internal.OrderRepository, productRepository internal.ProductRepository) *OrderUsecase {
	return &OrderUsecase{orderRepository, productRepository}
}

func (ou OrderUsecase) GetAllOrders() ([]*entity.Order, error) {
	orders, err := ou.orderRepository.GetAllOrders()
	if err != nil {
		log.Println(err.Error())
	}

	return orders, err
}

func (ou OrderUsecase) GetOrderItems(orderID int64) ([]*entity.OrderItem, error) {
	orderItems, err := ou.orderRepository.GetOrderItemsByID(orderID)
	if err != nil {
		log.Println(err.Error())
	}

	return orderItems, err
}

func (ou OrderUsecase) GetAnnualIncome() ([]*entity.AnnualIncome, error) {
	annualIncomes, err := ou.orderRepository.GetAnnualIncome()
	if err != nil {
		log.Println(err.Error())
	}

	return annualIncomes, err
}

func (ou OrderUsecase) GetDailyOrderCount() (int, error) {
	res, err := ou.orderRepository.GetDailyOrderCount()
	if err != nil {
		log.Println(err.Error())
	}

	return res, err
}

func (ou OrderUsecase) GetTotalOrderCount() (int, error) {
	res, err := ou.orderRepository.GetTotalOrderCount()
	if err != nil {
		log.Println(err.Error())
	}

	return res, err
}

func (ou OrderUsecase) GetLastDayIncome() (int, error) {
	res, err := ou.orderRepository.GetLastDayIncome()
	if err != nil {
		log.Println(err.Error())
	}

	return res, err
}

func (ou OrderUsecase) GetLastMonthIncome() (int, error) {
	res, err := ou.orderRepository.GetLastMonthIncome()
	if err != nil {
		log.Println(err.Error())
	}

	return res, err
}

func (ou OrderUsecase) Create(param entity.CreateOrderParam) (*entity.Order, error) {
	// check available quantity
	orderQuantity := map[int64]int{}
	productIDs := make([]int64, 0)

	for _, item := range param.Items {
		orderQuantity[item.ProductId] = item.Quantity
		productIDs = append(productIDs, item.ProductId)
	}

	products, err := ou.productRepository.GetProductsByIDs(productIDs...)
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

	order, err := ou.orderRepository.Create(param)
	if err != nil {
		return nil, err
	}

	return order, nil
}
