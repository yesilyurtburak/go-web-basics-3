package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/config"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/handlers"
)

// in place of built-in packages, we use a 3rd party package for routing: chi

func routes(app *config.AppConfig) http.Handler {
	// create a multiplexer to handle incoming http requests
	mux := chi.NewMux()

	// routing pages
	mux.Get("/", handlers.HomeHandler)
	mux.Get("/about", handlers.AboutHandler)
	return mux
}
