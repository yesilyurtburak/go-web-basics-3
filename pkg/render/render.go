package render

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/yesilyurtburak/go-web-basics-3/models"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/helpers"
)

// do not render everything while reloading the page, use templateCache instead.
var templateCache = make(map[string]*template.Template)

// this function creates a cache.
func makeTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gotmpl",
	}
	// this creates a new template from the given files
	tmpl, err := template.ParseFiles(templates...)
	helpers.ErrorCheck(err)
	templateCache[t] = tmpl // add the loaded page's template data to the templateCache map.
	return nil
}

// this function renders the template on the browser
func RenderTemplate(w http.ResponseWriter, t string, pd *models.PageData) {
	var tmpl *template.Template
	var err error
	// check if the template is already in cache
	_, inMap := templateCache[t]
	if !inMap {
		err = makeTemplateCache(t)
		helpers.ErrorCheck(err)
	} else {
		fmt.Println("Template is loaded from cache")
	}
	tmpl = templateCache[t]
	err = tmpl.Execute(w, pd) // writes the template to the response writer `w` by sending data `pd`
	helpers.ErrorCheck(err)
}
