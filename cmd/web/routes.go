package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrkouhadi/go-todolist/internal/config"
	"github.com/mrkouhadi/go-todolist/internal/handlers"
)

func Routes(app *config.Appconfiguration) http.Handler {
	// initialize the chi router
	mux := chi.NewRouter()

	// apply middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(Nosurf) //  ignore any POST request that doesn't have CSRF token
	mux.Use(LoadSession)
	// apply GET requests
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/alltodos", handlers.Repo.AllTodos)
	mux.Get("/search", handlers.Repo.SearchTodos)
	// apply POST requests

	// render files in the template(html)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// return the http.hanlder
	return mux
}
