package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pinnock/bmwawg/internal/config"
	"github.com/pinnock/bmwawg/internal/forms"
	"github.com/pinnock/bmwawg/internal/models"
	"github.com/pinnock/bmwawg/internal/render"
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

	render.RenderTemplate(w, r, "home.page.gohtml", &models.TemplateData{})
}

// About is the handler for the About page
func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	strData := map[string]string{}
	strData["test"] = "This is a test"
	strData["remote_ip"] = h.appCfg.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(
		w, r, "about.page.gohtml", &models.TemplateData{StringData: strData},
	)
}

func (h *Handlers) MontegoBaySuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(
		w, r, "montego-bay-suite.page.gohtml", &models.TemplateData{},
	)
}

func (h *Handlers) PortRoyalSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(
		w, r, "port-royal-suite.page.gohtml", &models.TemplateData{},
	)
}

func (h *Handlers) NegrilBungalow(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(
		w, r, "negril-bungalow.page.gohtml", &models.TemplateData{},
	)
}

func (h *Handlers) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(
		w, r, "search-availability.page.gohtml", &models.TemplateData{},
	)
}

type availabilityResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (h *Handlers) SearchAvailabilityJSON(
	w http.ResponseWriter, r *http.Request,
) {
	resp := availabilityResponse{
		OK:      true,
		Message: "Available!",
	}

	jsonResp, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (h *Handlers) PostSearchAvailability(
	w http.ResponseWriter, r *http.Request,
) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	fmt.Fprintf(w, "Start date is %s and end date is %s", start, end)
}

func (h *Handlers) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(
		w, r, "make-reservation.page.gohtml", &models.TemplateData{
			FormInfo: forms.New(nil),
		},
	)
}

func (h *Handlers) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	info := forms.New(r.PostForm)
	info.HasValue("first-name", r)
	if !info.Valid() {
		data := map[string]interface{}{}
		data["reservation"] = reservation

		render.RenderTemplate(
			w, r, "make-reservation.page.gohtml", &models.TemplateData{
				FormInfo:    info,
				GenericData: data,
			},
		)
		return
	}
}

func (h *Handlers) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.gohtml", &models.TemplateData{})
}
