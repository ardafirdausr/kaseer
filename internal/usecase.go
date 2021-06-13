package internal

import "github.com/ardafirdausr/go-pos/internal/entity"

type UserUsecase interface {
	GetUserByID(ID int64) (*entity.User, error)
	GetUserByCredential(email string, password string) (*entity.User, error)
	UpdateUser(ID int64, param entity.UpdateUserParam) (bool, error)
	UpdateUserPassword(ID int64, password string) (bool, error)
}

type ProductUsecase interface {
	GetAllProducts() ([]*entity.Product, error)
	GetProductByID(ID int64) (*entity.Product, error)
	GetProductByCode(code string) (*entity.Product, error)
	GetBestSellerProducts() ([]*entity.ProductSale, error)
	CreateProduct(param entity.CreateProductParam) (*entity.Product, error)
	UpdateProduct(ID int64, param entity.UpdateProductParam) (bool, error)
	DeleteProduct(ID int64) (bool, error)
}

type OrderUsecase interface {
	GetAllOrders() ([]*entity.Order, error)
	GetOrderItems(orderID int64) ([]*entity.OrderItem, error)
	GetAnnualIncome() ([]*entity.AnnualIncome, error)
	GetDailyOrderCount() (int, error)
	GetTotalOrderCount() (int, error)
	GetLastDayIncome() (int, error)
	GetLastMonthIncome() (int, error)
	Create(entity.CreateOrderParam) ([]*entity.Order, error)
}
