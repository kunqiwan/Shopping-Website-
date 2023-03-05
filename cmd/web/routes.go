package main

import (
	"github.com/KQW/my_page/pkg/config"
	"github.com/KQW/my_page/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// using the pat package to route incoming HTTP requests to specific handler functions
// create a custom http.Handler object to control over how requests are handled
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//
	//log the panic
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutPage))
	return mux
}
