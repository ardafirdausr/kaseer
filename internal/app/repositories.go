package app

import (
	"database/sql"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/repository/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type repositories struct {
	UserRepository    internal.UserRepository
	ProductRepository internal.ProductRepository
	OrderRepository   internal.OrderRepository
	UnitOfWork        internal.UnitOfWork
}

func newMySQLRepositories(DB *sql.DB) *repositories {
	return &repositories{
		UserRepository:    mysql.NewUserRepository(DB),
		ProductRepository: mysql.NewProductRepository(DB),
		OrderRepository:   mysql.NewOrderRepository(DB),
		UnitOfWork:        mysql.NewMySQLUnitOfWork(DB),
	}
}
