package handlers

import (
	"github.com/KQW/my_page/pkg/config"
	"github.com/KQW/my_page/pkg/render"
	"net/http"
)

// Repo use other class to update the Repo here
var Repo *Repository

// Repository Model is a software design mode,separate the logic that retrieves the data and maps it to the entity model
type Repository struct {
	App *config.AppConfig
}

// NewRepository creating a new instance of Repository and setting its App field to the value of the a parameter
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}
func (m *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
