package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func PreMiddlewares() []echo.MiddlewareFunc {
	middlewares := []echo.MiddlewareFunc{}
	middlewares = append(middlewares, middleware.RemoveTrailingSlash())
	// add more pre middlewares...
	return middlewares
}
