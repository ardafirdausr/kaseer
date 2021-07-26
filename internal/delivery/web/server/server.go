package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/ardafirdausr/kaseer/internal/delivery/web/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.Server.WriteTimeout = time.Minute * 2
	e.Server.ReadTimeout = time.Minute * 2

	debug := os.Getenv("DEBUG")
	isDebuging, err := strconv.ParseBool(debug)
	if err != nil {
		isDebuging = false
	}
	e.Debug = isDebuging

	renderer := NewHtmlRenderer()
	e.Renderer = renderer

	validator := &CustomValidator{validator: validator.New()}
	e.Validator = validator

	SentryDsn := os.Getenv("SENTRY_DSN")
	sentryMiddleware := middleware.Sentry(SentryDsn, isDebuging)
	e.Use(sentryMiddleware)

	errorHandler := &CustomHTTPErrorHandler{debug: isDebuging, logger: e.Logger}
	// errorHandler := &CustomHTTPErrorHandler{debug: isDebuging}
	e.HTTPErrorHandler = errorHandler.Handler

	preMiddlewares := middleware.PreMiddlewares()
	e.Pre(preMiddlewares...)

	sessionKey := os.Getenv("SESSION_KEY")
	store := sessions.NewCookieStore([]byte(sessionKey))
	e.Use(middleware.Session(store))

	e.Use(middleware.EchoLogger())
	e.Use(middleware.Recover())

	return e
}

func Start(e *echo.Echo) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	// Start server
	go func() {
		e.Logger.Info("Starting server...")
		if err := e.Start(host + ":" + port); err != nil {
			e.Logger.Info("Shutting down the server. error: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
