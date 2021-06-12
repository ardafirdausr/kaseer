package server

import (
	"log"
	"net/http"

	"github.com/ardafirdausr/go-pos/internal/entity"
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
		if ev.Message == "" {
			he.Message = http.StatusText(http.StatusBadRequest)
		} else {
			he.Message = ev.Message
		}
	} else if ent, ok := err.(*entity.ErrNotFound); ok {
		he.Code = http.StatusNotFound
		if ent.Message == "" {
			he.Message = http.StatusText(http.StatusNotFound)
		} else {
			he.Message = ent.Message
		}
	}

	if hub := sentryecho.GetHubFromContext(c); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			hub.CaptureException(err)
		})
		// if he.Code == http.StatusInternalServerError {
		// }
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
				c.Render(http.StatusNotFound, "500", nil)
			}
		}

		if err != nil {
			log.Println(err)
			// che.logger.Error(err)
		}
	}
}
