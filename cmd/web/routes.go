package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/config"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/handlers"
)

// in place of built-in packages, we use a 3rd party package for routing: chi

func routes(app *config.AppConfig) http.Handler {
	// create a multiplexer to handle incoming http requests
	mux := chi.NewMux()

	// using middlewares
	mux.Use(middleware.Recoverer) // built-in middleware for handling the panic gracefully.
	mux.Use(LogRequestInfo)       // custom middleware that is developed for a spesific purpose.

	// routing pages
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	return mux
}
