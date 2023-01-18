package handlers

import (
	"net/http"

	"github.com/yesilyurtburak/go-web-basics-3/models"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/config"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/render"
)

// Type definition for Repository pattern
type Repository struct {
	App *config.AppConfig
}

// Variable declaration for Repository pattern
var Repo *Repository

// Function definition for creating a new Repository
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// Function definition to handle routing with Repository pattern
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// set up an example cookie
	// whenever user goes to the homepage, they'are going to automatically have the userid placed inside of here as if they were entered it.
	m.App.Session.Put(r.Context(), "userid", "burakyesilyurt")

	render.RenderTemplate(w, "home.page.gotmpl", &models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["content"] = "We like to talk about ourselves here"

	// get `userid` from the session.
	userid := m.App.Session.GetString(r.Context(), "userid")
	strMap["userid"] = userid // add it to the strMap to send to the page template.

	render.RenderTemplate(w, "about.page.gotmpl", &models.PageData{StrMap: strMap})
	// created a strMap and send some information to about.page.gotmpl template via models.PageData
}
