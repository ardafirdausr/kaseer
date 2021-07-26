package internal

import (
	"context"
	"mime/multipart"

	"github.com/ardafirdausr/kaseer/internal/entity"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, ID int64) (*entity.User, error)
	GetUserByCredential(ctx context.Context, credential entity.UserCredential) (*entity.User, error)
	SaveUserPhoto(ctx context.Context, user *entity.User, file *multipart.FileHeader) (string, error)
	UpdateUser(ctx context.Context, ID int64, param entity.UpdateUserParam) (bool, error)
	UpdateUserPassword(ctx context.Context, ID int64, password string) (bool, error)
}

type ProductUsecase interface {
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
	GetProductByID(ctx context.Context, ID int64) (*entity.Product, error)
	GetProductByCode(ctx context.Context, code string) (*entity.Product, error)
	GetBestSellerProducts(ctx context.Context) ([]*entity.ProductSale, error)
	CreateProduct(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error)
	UpdateProduct(ctx context.Context, ID int64, param entity.UpdateProductParam) (bool, error)
	DeleteProduct(ctx context.Context, ID int64) (bool, error)
}

type OrderUsecase interface {
	GetAllOrders(ctx context.Context) ([]*entity.Order, error)
	GetOrderItems(ctx context.Context, orderID int64) ([]*entity.OrderItem, error)
	GetAnnualIncome(ctx context.Context) ([]*entity.AnnualIncome, error)
	GetDailyOrderCount(ctx context.Context) (int, error)
	GetTotalOrderCount(ctx context.Context) (int, error)
	GetLastDayIncome(ctx context.Context) (int, error)
	GetLastMonthIncome(ctx context.Context) (int, error)
	Create(ctx context.Context, param entity.CreateOrderParam) (*entity.Order, error)
}
