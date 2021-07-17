package main

import (
	"log"

	"github.com/ardafirdausr/kaseer/internal/app"
	"github.com/ardafirdausr/kaseer/internal/delivery/web"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalf("Failed initiate the app\n%v", err)
	}

	web.Start(app)
}
