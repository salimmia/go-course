package handlers

import (
	"net/http"

	"github.com/salimmia/go-course/pkg/render"
)

// Home is the Home page Handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page Handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
