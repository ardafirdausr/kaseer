package app

import (
	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/usecase"
)

type Usecases struct {
	UserUsecase    internal.UserUsecase
	ProductUsecase internal.ProductUsecase
	OrderUsecase   internal.OrderUsecase
}

func newUsecases(repos *repositories) *Usecases {
	userUsecase := usecase.NewUserUsecase(repos.UserRepository)
	productUsecase := usecase.NewProductUsecase(repos.ProductRepository)
	orderUsecase := usecase.NewOrderUsecase(repos.OrderRepository, repos.ProductRepository)
	return &Usecases{
		UserUsecase:    userUsecase,
		ProductUsecase: productUsecase,
		OrderUsecase:   orderUsecase,
	}
}
