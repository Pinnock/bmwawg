package handlers

import (
	"net/http"

	"github.com/pinnock/bmwawg/pkg/render"
)

// Home is the handler for the Home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

// About is the handler for the About page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}
