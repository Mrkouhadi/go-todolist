package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mrkouhadi/go-todolist/internal/config"
	"github.com/mrkouhadi/go-todolist/internal/forms"
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
	// get "greet" from the session
	val := repo.AppConfig.Session.GetString(r.Context(), "greet")
	data := make(map[string]interface{})
	data["PageTitle"] = "All todos Listed here !"
	data["PageDescription"] = "Description of our todos list page is here"
	data["PageKeywords"] = "Todos,todolist,hobbies,education,sport"
	data["greet"] = val
	rendertemplates.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// ******************* TODOS
// GET json
func (repo *Repository) AllTodos(w http.ResponseWriter, r *http.Request) {
	// stroe "greet" in to the session
	repo.AppConfig.Session.Put(r.Context(), "greet", "Hello the world of golang devs")

	todos := Mytodos // it's an array of dummy data already created down this file
	data := make(map[string]interface{})
	data["todos"] = todos
	data["PageTitle"] = "All todos Listed here !"
	data["PageDescription"] = "Description of our todos list page is here"
	data["PageKeywords"] = "Todos,todolist,hobbies,education,sport"
	rendertemplates.RenderTemplate(w, r, "alltodos.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// json experimental
type jsonRes struct {
	OK    bool          `json:"ok"`
	TODOS []models.TODO `json:"todos"`
}

func (repo *Repository) AllTodosjson(w http.ResponseWriter, r *http.Request) {
	resp := jsonRes{
		OK:    true,
		TODOS: Mytodos, // it's an array of dummy data already created down this file
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}
	// specify the content-type
	w.Header().Set("Content-Type", "application/json")
	// send data as json to the client
	w.Write(out)
}

// ******************* SEARCH
// GET
func (repo *Repository) SearchTodos(w http.ResponseWriter, r *http.Request) {

	// remove "greet" key/value from  the session
	repo.AppConfig.Session.Remove(r.Context(), "greet")
	data := make(map[string]interface{})
	data["PageTitle"] = "Searching for a new todo"
	data["PageDescription"] = "Description of our todos list page is here"
	data["PageKeywords"] = "Todos,todolist,hobbies,education,sport"
	rendertemplates.RenderTemplate(w, r, "search.page.tmpl", &models.TemplateData{
		Data: data,
	})
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
}

// ****************** ADD NEW TODOS
//
//	GET
func (repo *Repository) AddNewTodo(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["PageTitle"] = "Add a new Todo"
	data["PageDescription"] = "Description of our todos list page is here"
	data["PageKeywords"] = "Todos,todolist,hobbies,education,sport"
	var emptyTodo models.TODO
	data["todo"] = emptyTodo
	rendertemplates.RenderTemplate(w, r, "addnewtodo.page.tmpl", &models.TemplateData{
		Data: data,
		Form: forms.New(nil),
	})
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
		Time:        Substring(time.Now().String(), 0, 10),
	}
	// form
	form := forms.New(r.PostForm)

	form.Required("title", "email", "content")
	form.MinLength("title", 5, r)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["todo"] = newTodo
		rendertemplates.RenderTemplate(w, r, "addnewtodo.page.tmpl", &models.TemplateData{
			Data: data,
			Form: form,
		})
		return
	}
	// add newTodo to the array of dummy data down this file
	Mytodos = append(Mytodos, newTodo)
	// post the newly posted todo in a session
	repo.AppConfig.Session.Put(r.Context(), "newTodo", newTodo)
	// redirect the user to /alltodos after successfully submitting the data
	http.Redirect(w, r, "/todo-review", http.StatusSeeOther)
}

func (repo *Repository) TodoReview(w http.ResponseWriter, r *http.Request) {
	newtodo, ok := repo.AppConfig.Session.Get(r.Context(), "newTodo").(models.TODO)
	if !ok {
		log.Println("Cannot get data from session!")
		repo.AppConfig.Session.Put(r.Context(), "error", "Cannot get data from session!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]interface{})
	data["PageTitle"] = "Review of the newly added Todo"
	data["PageDescription"] = "Description of our todos list page is here"
	data["PageKeywords"] = "Todos,todolist,hobbies,education,sport"
	data["newTodo"] = newtodo
	repo.AppConfig.Session.Remove(r.Context(), "newTodo")
	rendertemplates.RenderTemplate(w, r, "todo-review.page.tmpl", &models.TemplateData{
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

// /// testing DATA
var Mytodos []models.TODO = []models.TODO{
	{
		Title:       "hello gophers",
		Email:       "bryan@kouhadi.com",
		Content:     "this is  the content of my todo ",
		IsImportant: true,
		IsDone:      false,
		Time:        Substring(time.Now().String(), 0, 10),
	},
	{
		Title:       "hello gophers 2",
		Email:       "bryan@kouhadiii.com",
		Content:     "this is  the content of my todo 2",
		IsImportant: false,
		IsDone:      true,
		Time:        Substring(time.Now().String(), 0, 10),
	},
}

// get substring
func Substring(str string, start, end int) string {
	return strings.TrimSpace(str[start:end])
}
