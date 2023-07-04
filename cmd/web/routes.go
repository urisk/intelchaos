package main

import (
	"github.com/go-chi/chi/v5"
	"intelchaos/pkg/config"
	"intelchaos/pkg/handlers"
	"net/http"
)

func routes (app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
