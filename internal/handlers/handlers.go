package handlers

import (
	"log"
	"net/http"
	"time"

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
	rendertemplates.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// TODOS
func (repo *Repository) AllTodos(w http.ResponseWriter, r *http.Request) {
	// greet := repo.AppConfig.Session.GetString(r.Context(), "greet")
	rendertemplates.RenderTemplate(w, r, "alltodos.page.tmpl", &models.TemplateData{})
}

// SEARCH
func (repo *Repository) SearchTodos(w http.ResponseWriter, r *http.Request) {
	rendertemplates.RenderTemplate(w, r, "search.page.tmpl", &models.TemplateData{})
}

// add new todo
//
//	GET
func (repo *Repository) AddNewTodo(w http.ResponseWriter, r *http.Request) {
	rendertemplates.RenderTemplate(w, r, "addnewtodo.page.tmpl", &models.TemplateData{})
}

// POST
func (repo *Repository) PostNewTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	// data
	newTodo := models.TODO{
		Title:       r.Form.Get("title"),
		Email:       r.Form.Get("email"),
		Content:     r.Form.Get("content"),
		IsImportant: convertStrToBool(r.Form.Get("important")),
		IsDone:      false,
		Time:        time.Now(),
	}

	data := make(map[string]interface{})
	data["todo"] = newTodo

	log.Println(newTodo)
	rendertemplates.RenderTemplate(w, r, "addnewtodo.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// convertStrToBool helps us convert value of checkbox from "on"/"off" to true/false
func convertStrToBool(s string) bool {
	return s == "on"
}
