package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	repositories *repositories
	Usecases     *Usecases
}

func New() (*App, error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	MySQLURI := os.Getenv("MYSQL_URI")
	MySQL, err := connectToMySQL(MySQLURI)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	app := &App{}
	app.repositories = newMySQLRepositories(MySQL)
	app.Usecases = newUsecases(app.repositories)
	return app, nil
}
