package handlers

import (
	"net/http"

	"github.com/mrkouhadi/go-todolist/internal/config"
	"github.com/mrkouhadi/go-todolist/internal/models"
	"github.com/mrkouhadi/go-todolist/internal/rendertemplates"
)

/******************************************** REPOSITORY ****************************************/

type Repository struct {
	AppConfig *config.Appconfiguration
}

var Repo *Repository

func NewRepository(AC *config.Appconfiguration) *Repository {
	return &Repository{
		AppConfig: AC,
	}
}

/******************************************** handlers ****************************************/
func NewHandlers(R *Repository) {
	Repo = R
}

// HOME
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// repo.AppConfig.Session.Put(r.Context(), "greet", "Hello the world of golang devs")
	// w.Write([]byte("hello home page"))
	rendertemplates.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// TODOS
func (repo *Repository) AllTodos(w http.ResponseWriter, r *http.Request) {
	// greet := repo.AppConfig.Session.GetString(r.Context(), "greet")
	// w.Write([]byte(greet))
	rendertemplates.RenderTemplate(w, r, "alltodos.page.tmpl", &models.TemplateData{})
}
