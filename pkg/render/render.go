package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	if err != nil {
		log.Fatal(err)
	}
	templateCache[t] = tmpl // add the loaded page's template data to the templateCache map.
	return nil
}

// this function renders the template on the browser
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	// check if the template is already in cache
	_, inMap := templateCache[t]
	if !inMap {
		err = makeTemplateCache(t)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Template is loaded from cache")
	}
	tmpl = templateCache[t]
	err = tmpl.Execute(w, nil) // writes the template to the response writer `w`
	if err != nil {
		log.Fatal(err)
	}
}
