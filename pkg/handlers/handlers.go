package handlers

import (
	"net/http"

	"github.com/salimmia/go-course/pkg/config"
	"github.com/salimmia/go-course/pkg/render"
)

// Repo is the repository by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates new Repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the Handlers
func NewHandler(r *Repository){
	Repo = r
}

// Home is the Home page Handler
func (m *Repository)Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page Handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
