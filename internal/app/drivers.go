package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/ardafirdausr/kaseer/internal/driver"
)

type drivers struct {
	MySQL *sql.DB
}

func newDrivers() (*drivers, error) {
	MySQLURI := os.Getenv("MYSQL_URI")
	MySQLConn, err := driver.ConnectToMySQLDB(MySQLURI)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	drivers := new(drivers)
	drivers.MySQL = MySQLConn
	return drivers, nil
}
