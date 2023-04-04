package main

import (
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
