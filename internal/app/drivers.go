package app

import (
	"database/sql"
	"log"

	"github.com/ardafirdausr/kaseer/internal/driver"
)

type drivers struct {
	MySQL *sql.DB
}

func newDrivers() (*drivers, error) {
	MySQLConn, err := driver.ConnectToMySQLDB()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	drivers := new(drivers)
	drivers.MySQL = MySQLConn
	return drivers, nil
}
