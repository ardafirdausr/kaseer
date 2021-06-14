package app

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
	queries := []string{"parseTime=true"}
	queryString := strings.Join(queries, "&")

	DBURIx := strings.TrimRight(DBURI, "/")
	MySQLURI := fmt.Sprintf("%s?%s", DBURIx, queryString)
	DB, err := sql.Open("mysql", MySQLURI)
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
