package app

import (
	"database/sql"
	"log"
	"net/url"

	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/repository/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type repositories struct {
	UserRepository    internal.UserRepository
	ProductRepository internal.ProductRepository
	OrderRepository   internal.OrderRepository
}

func newMySQLRepositories(DB *sql.DB) *repositories {
	return &repositories{
		UserRepository:    mysql.NewUserRepository(DB),
		ProductRepository: mysql.NewProductRepository(DB),
		OrderRepository:   mysql.NewOrderRepository(DB),
	}
}

func connectToMySQL(DBURI string) (*sql.DB, error) {
	u, err := url.Parse(DBURI)
	if err != nil {
		log.Fatal(err.Error())
	}

	q := u.Query()
	q.Set("parseTime", "true")
	u.RawQuery = q.Encode()

	DBURI = u.String()
	DB, err := sql.Open("mysql", DBURI)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return DB, nil
}
