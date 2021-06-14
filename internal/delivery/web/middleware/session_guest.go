package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SessionGuest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("GO-POS", c)
			if err != nil {
				return err
			}

			if sess.IsNew {
				sess.Options = &sessions.Options{
					Path:     "/",
					MaxAge:   86400 * 7,
					HttpOnly: true,
				}
			}

			if _, ok := sess.Values["user"]; ok {
				return c.Redirect(http.StatusSeeOther, "/dashboard")
			}

			return next(c)
		}
	}
}
