package main

import (
	"log"
	"net/http"

	"github.com/pinnock/bmwawg/pkg/config"
	"github.com/pinnock/bmwawg/pkg/handlers"
	"github.com/pinnock/bmwawg/pkg/render"
)

const port string = ":8080"

func main() {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("failed to create template cache: %v\n", err)
	}
	cfg := &config.AppConfig{
		TemplateCache: tc,
		UseCache:      false,
	}
	render.SetConfig(cfg)

	h := handlers.NewHandlers(cfg)
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/about", h.About)

	log.Printf("Starting web server on port %s\n", port)
	http.ListenAndServe(port, nil)
}
