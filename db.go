package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(host string, port string, username string, password string, DBName string) (*sql.DB, error) {
	DBURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, DBName)
	DB, err := sql.Open("mysql", DBURI)
	if err != nil {
		return nil, err
	}

	return DB, nil
}
