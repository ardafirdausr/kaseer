package internal

import "github.com/ardafirdausr/go-pos/internal/entity"

type UserRepository interface {
	GetUserByID(ID int64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateByID(ID int64, param entity.UpdateUserParam) (bool, error)
	UpdatePasswordByID(ID int64, password string) (bool, error)
}

type ProductRepository interface {
	GetAllProducts() ([]*entity.Product, error)
	GetBestSellerProducts() ([]*entity.ProductSale, error)
	GetProductByCode(code string) (*entity.Product, error)
	GetProductByID(ID int64) (*entity.Product, error)
	Create(param entity.CreateProductParam) (*entity.Product, error)
	UpdateByID(ID int64, param entity.UpdateProductParam) (bool, error)
	DeleteByID(ID int64) (bool, error)
}

type OrderRepository interface {
	GetAllOrders() ([]*entity.Order, error)
	GetAnnualIncome() ([]*entity.AnnualIncome, error)
	GetDailyOrderCount() (int, error)
	GetTotalOrderCount() (int, error)
	GetLastDayIncome() (int, error)
	GetLastMonthIncome() (int, error)
	GetOrderItemsByID(ID int64) ([]*entity.OrderItem, error)
	Create(param entity.CreateOrderParam) (*entity.Order, error)
}
