package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pinnock/bmwawg/pkg/config"
	"github.com/pinnock/bmwawg/pkg/handlers"
)

func routes(c *config.AppConfig) http.Handler {
	h := handlers.NewHandlers(c)

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(LogRequest)
	mux.Use(NoSurf)
	mux.Get("/", http.HandlerFunc(h.Home))
	mux.Get("/about", http.HandlerFunc(h.About))
	return mux
}
