package rendertemplates

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/mrkouhadi/go-todolist/internal/config"
	"github.com/mrkouhadi/go-todolist/internal/models"
)

func AddDefaultData(tmplData *models.TemplateData, r *http.Request) *models.TemplateData {
	tmplData.Flash = app.Session.PopString(r.Context(), "flash")
	tmplData.Error = app.Session.PopString(r.Context(), "error")
	tmplData.Warning = app.Session.PopString(r.Context(), "warning")

	tmplData.CSRFToken = nosurf.Token(r)
	return tmplData
}

// render tamplates
var app *config.Appconfiguration

// newTemplates sets the config for the template package
func NewTemplates(a *config.Appconfiguration) {
	app = a
}

// RenderTemplate renders the requested template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, tmplData *models.TemplateData) {
	// Get the template cache from the AppConfig
	var tmplCache map[string]*template.Template
	if app.UseCache {
		tmplCache = app.TemplateCache
	} else {
		tmplCache, _ = CreateTemplateCache()
	}

	// get requested template from cached templates
	t, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("Could not get the template from Cached templates ! ")
	}
	buf := new(bytes.Buffer)

	tmplData = AddDefaultData(tmplData, r)
	_ = t.Execute(buf, tmplData)

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache create cache for templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		fileName := filepath.Base(page) // from "templates/home.page.tmpl" to "home.page.tmpl"
		templSet, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// look for any layout that exist in that directory
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templSet, err = templSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[fileName] = templSet
	}
	return myCache, nil
}
