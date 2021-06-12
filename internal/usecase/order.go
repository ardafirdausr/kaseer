package usecase

import (
	"log"

	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/entity"
)

type OrderUsecase struct {
	orderRepository internal.OrderRepository
}

func NewOrderUsecase(orderRepository internal.OrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepository: orderRepository}
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

func (ou OrderUsecase) Create(entity.CreateOrderParam) ([]*entity.Order, error) {
	return nil, nil
}
