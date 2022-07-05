package main

import (
	"html/template"
	"log"
	"net/http"
)

const port string = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Starting web server on port %s\n", port)
	http.ListenAndServe(port, nil)
}

// Home is the handler for the Home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.gohtml")
}

// About is the handler for the About page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.gohtml")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}
	parsedTemplate.Execute(w, nil)
}
