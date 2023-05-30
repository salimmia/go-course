package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/salimmia/go-course/pkg/config"
	"github.com/salimmia/go-course/pkg/handlers"
	"github.com/salimmia/go-course/pkg/render"
	// "os"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	// initializing Template cache
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// store template cache to app config
	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/home", handlers.Repo.Home)

	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
