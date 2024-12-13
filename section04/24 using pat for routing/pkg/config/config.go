package config

import (
	"text/template"
)

/* config is imported by other parts of the app
but it doesn't import  anything else from the app itself */

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	// InfoLog       *log.Logger
}
