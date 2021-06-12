package controller

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func renderPage(c echo.Context, page string, title string, additionalData echo.Map) error {
	session, _ := session.Get("GO-POS", c)
	success := true
	msg := ""

	errs := session.Flashes("error")
	if len(errs) > 0 {
		success = false
		ems := session.Flashes("error_message")
		if len(ems) > 0 {
			msg = ems[0].(string)
		}
	} else {
		sms := session.Flashes("success_message")
		if len(sms) > 0 {
			msg = sms[0].(string)
		}
	}

	session.Save(c.Request(), c.Response())

	data := echo.Map{
		"Title":   title,
		"URL":     c.Request().URL,
		"Success": success,
		"Message": msg,
		"Data":    additionalData,
	}

	return c.Render(http.StatusOK, page, data)
}

func json(c echo.Context, code int, message string, data interface{}) error {
	payload := map[string]interface{}{
		"message": message,
		"data":    data,
	}
	return c.JSON(code, payload)
}
