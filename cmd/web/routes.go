package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/pinnock/bmwawg/pkg/config"
	"github.com/pinnock/bmwawg/pkg/handlers"
)

func routes(c *config.AppConfig) http.Handler {
	mux := pat.New()
	h := handlers.NewHandlers(c)
	mux.Get("/", http.HandlerFunc(h.Home))
	mux.Get("/about", http.HandlerFunc(h.About))
	return mux
}
