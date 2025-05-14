package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQLDB() (*sql.DB, error) {
	retryCount := 0
	retryLimit := 3
	retryInterval := 3
	var db *sql.DB
	for retryCount < retryLimit {
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}

		mysqlConfig, err := mysql.NewConnector(&mysql.Config{
			Addr:         fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
			User:         os.Getenv("MYSQL_USER"),
			Passwd:       os.Getenv("MYSQL_PASS"),
			DBName:       os.Getenv("MYSQL_DATABASE"),
			ParseTime:    true,
			Net:          "tcp",
			Loc:          loc,
			Collation:    "utf8mb4_general_ci",
			Timeout:      10 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		})
		if err != nil {
			log.Printf("Failed to connect to MySQL database. Retrying in %d seconds...\n", retryInterval)
			retryCount++
			if retryCount >= retryLimit {
				log.Fatal(err.Error())
				return nil, err
			}

			time.Sleep(time.Duration(retryInterval) * time.Second)
			retryInterval *= 2
			continue
		}

		// db, err := sql.Open("mysql", "mysql://root:Rahasia123@tcp(mysql:3306)/kaseer?parseTime=true")
		db = sql.OpenDB(mysqlConfig)
		db.SetConnMaxIdleTime(30)
		db.SetConnMaxLifetime(300)
		db.SetMaxIdleConns(3)
		db.SetMaxOpenConns(5)
		if err = db.Ping(); err != nil {
			log.Fatal(err.Error())
			return nil, err
		}

		break
	}

	return db, nil
}
