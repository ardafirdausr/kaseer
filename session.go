package main

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

func NewSessionStore(sessionKey string) *sessions.CookieStore {
	sessionStore := sessions.NewCookieStore([]byte(sessionKey))
	sessionStore.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   60 * 60 * 2,
		HttpOnly: true,
	}

	gob.Register(&User{})

	return sessionStore
}