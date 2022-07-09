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
	remoteIP := r.RemoteAddr
	h.appCfg.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the handler for the About page
func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	strData := map[string]string{}
	strData["test"] = "This is a test"
	strData["remote_ip"] = h.appCfg.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(
		w,
		"about.page.gohtml",
		&models.TemplateData{StringData: strData},
	)
}

func (h *Handlers) MontegoBaySuite(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "montego-bay-suite.page.gohtml", &models.TemplateData{})
}

func (h *Handlers) PortRoyalSuite(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "port-royal-suite.page.gohtml", &models.TemplateData{})
}

func (h *Handlers) NegrilBungalow(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "negril-bungalow.page.gohtml", &models.TemplateData{})
}

func (h *Handlers) SearchAvailability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search-availability.page.gohtml", &models.TemplateData{})
}

func (h *Handlers) MakeReservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "make-reservation.page.gohtml", &models.TemplateData{})
}

func (h *Handlers) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "contact.page.gohtml", &models.TemplateData{})
}
