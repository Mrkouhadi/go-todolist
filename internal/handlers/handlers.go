package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mrkouhadi/go-todolist/internal/config"
	"github.com/mrkouhadi/go-todolist/internal/models"
	"github.com/mrkouhadi/go-todolist/internal/rendertemplates"
)

/**************************************************************************
// REPOSITORY
***************************************************************************/

type Repository struct {
	AppConfig *config.Appconfiguration
}

var Repo *Repository

func NewRepository(AC *config.Appconfiguration) *Repository {
	return &Repository{
		AppConfig: AC,
	}
}

/*
***********************************************************************************
// handlers
***********************************************************************************
*/
func NewHandlers(R *Repository) {
	Repo = R
}

// ******************* HOME
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// repo.AppConfig.Session.Put(r.Context(), "greet", "Hello the world of golang devs")
	rendertemplates.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// ******************* TODOS
// GET json
type jsonRes struct {
	OK      bool          `json:"ok"`
	MESSAGE []models.TODO `json:"message"`
}

func (repo *Repository) AllTodos(w http.ResponseWriter, r *http.Request) {
	// greet := repo.AppConfig.Session.GetString(r.Context(), "greet")
	resp := jsonRes{
		OK: true,
		MESSAGE: []models.TODO{
			{
				Title:       "hello gophers",
				Email:       "bryan@kouhadi.com",
				Content:     "this is  the content of my todo ",
				IsImportant: true,
				IsDone:      false,
				Time:        time.Now(),
			},
			{
				Title:       "hello gophers 2",
				Email:       "bryan@kouhadi2.com",
				Content:     "this is  the content of my todo 2",
				IsImportant: false,
				IsDone:      true,
				Time:        time.Now(),
			},
		},
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}
	// specify the content-type
	w.Header().Set("Content-Type", "application/json")
	// send data as json to the client
	w.Write(out)
	// rendertemplates.RenderTemplate(w, r, "alltodos.page.tmpl", &models.TemplateData{})
}

// ******************* SEARCH
// GET
func (repo *Repository) SearchTodos(w http.ResponseWriter, r *http.Request) {
	rendertemplates.RenderTemplate(w, r, "search.page.tmpl", &models.TemplateData{})
}

// POST
func (repo *Repository) PostSearchTodos(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	title := r.Form.Get("title")
	done := convertStrToBool(r.Form.Get("done"))
	important := convertStrToBool(r.Form.Get("important"))

	log.Println(title, done, important)
	w.Write([]byte(fmt.Sprintf("title: %s", title)))

	// rendertemplates.RenderTemplate(w, r, "search.page.tmpl", &models.TemplateData{})
}

// ****************** ADD NEW TODOS
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

	rendertemplates.RenderTemplate(w, r, "addnewtodo.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

/********************************************************************************
//  HELPERS
**********************************************************************************/
// convertStrToBool helps us convert value of checkbox from "on"/"off" to true/false
func convertStrToBool(s string) bool {
	return s == "on"
}
