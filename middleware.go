package main

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := SessionStore.Get(r, SessionName)
		if session.IsNew {
			session.Save(r, w)
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		userId, isUserIdAvailable := session.Values["user_id"]
		if !isUserIdAvailable {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		user, err := findUserById(userId.(int64))
		if err != nil || user == nil {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		session.Values["user"] = user
		err = session.Save(r, w)
		if err != nil {
			log.Println(err.Error())
		}

		next.ServeHTTP(w, r)
	})
}
