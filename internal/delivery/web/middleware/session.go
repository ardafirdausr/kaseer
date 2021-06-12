package middleware

import (
	"encoding/gob"

	"github.com/ardafirdausr/go-pos/internal/entity"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Session(store sessions.Store) echo.MiddlewareFunc {
	gob.Register(&entity.User{})
	cfg := session.Config{Store: store}
	return session.MiddlewareWithConfig(cfg)
}
