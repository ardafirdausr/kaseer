package app

import (
	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/usecase"
)

type Usecases struct {
	UserUsecase    internal.UserUsecase
	ProductUsecase internal.ProductUsecase
	OrderUsecase   internal.OrderUsecase
}

func newUsecases(repos *repositories) *Usecases {
	userUsecase := usecase.NewUserUsecase(repos.UserRepository)
	productUsecase := usecase.NewProductUsecase(repos.ProductRepository)
	orderUsecase := usecase.NewOrderUsecase(repos.OrderRepository)
	return &Usecases{
		UserUsecase:    userUsecase,
		ProductUsecase: productUsecase,
		OrderUsecase:   orderUsecase,
	}
}
