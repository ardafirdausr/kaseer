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

func newUsecases(app *App) *Usecases {
	userUsecase := usecase.NewUserUsecase(app.repositories.UserRepository, app.services.Storage)
	productUsecase := usecase.NewProductUsecase(app.repositories.ProductRepository)
	orderUsecase := usecase.NewOrderUsecase(app.repositories.OrderRepository, app.repositories.ProductRepository, app.repositories.UnitOfWork)
	return &Usecases{
		UserUsecase:    userUsecase,
		ProductUsecase: productUsecase,
		OrderUsecase:   orderUsecase,
	}
}
