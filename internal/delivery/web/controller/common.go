package controller

import (
	"net/http"

	"github.com/ardafirdausr/go-pos/internal/entity"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func renderPage(c echo.Context, page string, title string, additionalData echo.Map) error {
	data := echo.Map{
		"Title":   title,
		"URL":     c.Request().URL,
		"Data":    additionalData,
		"Error":   echo.Map{},
		"Success": echo.Map{},
	}

	sess, _ := session.Get("GO-POS", c)
	ems := sess.Flashes("error_message")
	if len(ems) > 0 {
		errs := data["Error"].(echo.Map)
		errs["Message"] = ems[0].(string)
	}

	evl := sess.Flashes("error_validation")
	if len(evl) > 0 {
		if ev, ok := evl[0].(*entity.ErrValidation); ok {
			errs := data["Error"].(echo.Map)
			errs["Message"] = ev.Message
			errs["Errors"] = ev.Errors
		}
	}

	sms := sess.Flashes("success_message")
	if len(sms) > 0 {
		succ := data["Success"].(echo.Map)
		succ["Message"] = sms[0].(string)
	}

	sess.Save(c.Request(), c.Response())
	return c.Render(http.StatusOK, page, data)
}

func responseJson(c echo.Context, code int, message string, data interface{}) error {
	payload := map[string]interface{}{
		"message": message,
		"data":    data,
	}
	return c.JSON(code, payload)
}

func responseErrorJson(c echo.Context, code int, message string, errs interface{}) error {
	payload := map[string]interface{}{
		"message": message,
		"errors":  errs,
	}
	return c.JSON(code, payload)
}
