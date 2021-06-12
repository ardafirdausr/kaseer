package driver

import (
	"database/sql"
	"log"
	"net/url"
)

func ConnectToMySQLDB(DBURI string) (*sql.DB, error) {
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
