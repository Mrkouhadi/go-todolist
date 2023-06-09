package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type Appconfiguration struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
