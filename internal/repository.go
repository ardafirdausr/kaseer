package internal

import (
	"context"

	"github.com/ardafirdausr/kaseer/internal/entity"
)

type UnitOfWork interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type UserRepository interface {
	GetUserByID(ctx context.Context, ID int64) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateByID(ctx context.Context, ID int64, param entity.UpdateUserParam) (bool, error)
	UpdatePasswordByID(ctx context.Context, ID int64, password string) (bool, error)
}

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
	GetBestSellerProducts(ctx context.Context) ([]*entity.ProductSale, error)
	GetProductsByIDs(ctx context.Context, IDs ...int64) ([]*entity.Product, error)
	GetProductByCode(ctx context.Context, code string) (*entity.Product, error)
	GetProductByID(ctx context.Context, ID int64) (*entity.Product, error)
	Create(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error)
	UpdateByID(ctx context.Context, ID int64, param entity.UpdateProductParam) (bool, error)
	DecrementProductByIDs(ctx context.Context, IDDecrementMap map[int64]int) error
	DeleteByID(ctx context.Context, ID int64) (bool, error)
}

type OrderRepository interface {
	GetAllOrders(ctx context.Context) ([]*entity.Order, error)
	GetAnnualIncome(ctx context.Context) ([]*entity.AnnualIncome, error)
	GetDailyOrderCount(ctx context.Context) (int, error)
	GetTotalOrderCount(ctx context.Context) (int, error)
	GetLastDayIncome(ctx context.Context) (int, error)
	GetLastMonthIncome(ctx context.Context) (int, error)
	GetOrderItemsByID(ctx context.Context, ID int64) ([]*entity.OrderItem, error)
	Create(ctx context.Context, param entity.CreateOrderParam) (*entity.Order, error)
	CreateOrderItems(ctx context.Context, orderId int64, items []*entity.CreateOrderItemParam) error
}
