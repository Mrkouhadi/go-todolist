package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// csrf : ignore any POST request that doesn't have CSRF token
func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// sessions loader
func LoadSession(next http.Handler) http.Handler {
	return todoSessionManager.LoadAndSave(next)
}

// experimental middleware that only log "I am comming from Middleware !" to the console
func LogMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am comming from Middleware !")
		next.ServeHTTP(w, r)
	})
}
