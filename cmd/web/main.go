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

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/home", handlers.Home)

	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
