package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}
	parsedTemplate.Execute(w, nil)
}
