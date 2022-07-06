package handlers

import (
	"net/http"

	"github.com/pinnock/bmwawg/pkg/config"
	"github.com/pinnock/bmwawg/pkg/models"
	"github.com/pinnock/bmwawg/pkg/render"
)

type Handlers struct {
	appCfg *config.AppConfig
}

func NewHandlers(c *config.AppConfig) *Handlers {
	return &Handlers{
		appCfg: c,
	}
}

// Home is the handler for the Home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the handler for the About page
func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	strData := map[string]string{}
	strData["test"] = "This is a test"
	render.RenderTemplate(
		w,
		"about.page.gohtml",
		&models.TemplateData{StringData: strData},
	)
}
