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
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoadAndSave)

	mux.Get("/", http.HandlerFunc(h.Home))
	mux.Get("/about", http.HandlerFunc(h.About))
	mux.Get("/montego-bay-suite", http.HandlerFunc(h.MontegoBaySuite))
	mux.Get("/port-royal-suite", http.HandlerFunc(h.PortRoyalSuite))
	mux.Get("/negril-bungalow", http.HandlerFunc(h.NegrilBungalow))
	mux.Get("/search-availability", http.HandlerFunc(h.SearchAvailability))
	mux.Post("/search-availability", http.HandlerFunc(h.PostSearchAvailability))
	mux.Get("/make-reservation", http.HandlerFunc(h.MakeReservation))
	mux.Get("/contact", http.HandlerFunc(h.Contact))

	serveStaticFiles := http.StripPrefix(
		"/static",
		http.FileServer(http.Dir("./static/")),
	)

	mux.Handle("/static/*", serveStaticFiles)
	return mux
}
