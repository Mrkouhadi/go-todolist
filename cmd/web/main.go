package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrkouhadi/go-todolist/internal/config"
	"github.com/mrkouhadi/go-todolist/internal/handlers"
	"github.com/mrkouhadi/go-todolist/internal/rendertemplates"
)

const portNumber = ":8080"

var appConfig config.Appconfiguration

// initialize the session
var todoSessionManager *scs.SessionManager

func main() {
	// turn production mode or not yet
	appConfig.InProduction = false

	// set up the session
	todoSessionManager = scs.New()
	todoSessionManager.Lifetime = 24 * time.Hour
	todoSessionManager.Cookie.Persist = true
	todoSessionManager.Cookie.SameSite = http.SameSiteLaxMode
	todoSessionManager.Cookie.Secure = appConfig.InProduction
	appConfig.Session = todoSessionManager

	// initialize the rendering of templates
	tmplCache, err := rendertemplates.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create templates cache !!!")
	}
	appConfig.TemplateCache = tmplCache
	appConfig.UseCache = false

	// initialize the Repository and create newHanlders
	newRepo := handlers.NewRepository(&appConfig)
	handlers.NewHandlers(newRepo)

	rendertemplates.NewTemplates(&appConfig)
	// run the server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&appConfig),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
