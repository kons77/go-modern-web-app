package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp19/pkg/config"
	"webapp19/pkg/handlers"
	"webapp19/pkg/render"
)

const portNumber = ":8080"

// main os the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app) // will give our app the render component or the render package of our app access to this app config variable

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	// enter localhost to a browser / 127.0.0.1:8080
	_ = http.ListenAndServe(portNumber, nil)
}
