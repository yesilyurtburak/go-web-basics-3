package handlers

import (
	"net/http"

	"github.com/yesilyurtburak/go-web-basics-3/pkg/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gotmpl")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gotmpl")
}
