package server

import (
	"log"
	"net/http"

	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

type CustomHTTPErrorHandler struct {
	debug  bool
	logger echo.Logger
}

func (che CustomHTTPErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if ev, ok := err.(*entity.ErrValidation); ok {
		he.Code = http.StatusBadRequest
		he.Message = ev.Message
	} else if ent, ok := err.(*entity.ErrNotFound); ok {
		he.Code = http.StatusNotFound
		he.Message = ent.Message
	}

	if he.Message == "" {
		he.Message = http.StatusText(he.Code)
	}

	if hub := sentryecho.GetHubFromContext(c); hub != nil {
		if he.Code == http.StatusInternalServerError {
			hub.WithScope(func(scope *sentry.Scope) {
				hub.CaptureException(err)
			})
		}
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			switch he.Code {
			case http.StatusUnauthorized:
				c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			case http.StatusNotFound:
				c.Render(http.StatusNotFound, "404", nil)
			case http.StatusInternalServerError:
				var data echo.Map
				if che.debug {
					data = echo.Map{
						"Success": false,
						"Message": err.Error(),
					}
				}
				c.Render(http.StatusNotFound, "500", data)
			}
		}

		if err != nil {
			log.Println(he.Message)
			log.Println(he.Error())
			che.logger.Error(err)
		}
	}
}
