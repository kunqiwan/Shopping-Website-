package config

import "html/template"

// AppConfig holds the application config,to decrease multiple run
// in render part, avoid creating template cache every time
// in handler part,it needs to access to all kinds of configuration setting
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
