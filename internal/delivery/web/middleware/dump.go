package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Dump() echo.MiddlewareFunc {
	handler := func(c echo.Context, reqBody []byte, resBody []byte) {
		request := fmt.Sprintf("%s %s", c.Request().Method, c.Request().URL.Path)
		req := string(reqBody)
		res := string(resBody)
		c.Logger().Debug(request, req, res)
	}
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: handler,
	})
}
