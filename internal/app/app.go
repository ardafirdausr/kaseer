package app

import (
	"log"

	"github.com/joho/godotenv"
)

type App struct {
	repositories *repositories
	drivers      *drivers
	services     *services
	Usecases     *Usecases
}

func New() (*App, error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	drivers, err := newDrivers()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	app := new(App)
	app.drivers = drivers
	app.repositories = newMySQLRepositories(app.drivers.MySQL)
	app.services = NewServices()
	app.Usecases = newUsecases(app)
	return app, nil
}
