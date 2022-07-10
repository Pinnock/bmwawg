package models

import "github.com/pinnock/bmwawg/internal/forms"

// TemplateData holds data sent from handler to template
type TemplateData struct {
	IntData     map[string]int
	StringData  map[string]string
	FloatData   map[string]float32
	GenericData map[string]interface{}
	CSRFToken   string
	Flash       string
	Warning     string
	Error       string
	FormInfo    *forms.Info
}
