package middleware

import (
	"log"

	"github.com/ardafirdausr/kaseer/internal/entity"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SessionAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("kaseer", c)
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

			user, ok := sess.Values["user"]
			if !ok {
				return echo.ErrUnauthorized
			}

			user, ok = user.(*entity.User)
			if !ok {
				log.Println("Failed to parse user session")
				return echo.ErrUnauthorized
			}

			c.Set("user", user)
			return next(c)
		}
	}
}
