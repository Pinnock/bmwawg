package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

// Home is the handler for the Home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page\n")
}

// Abou is the handler for the About page
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About page\n")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Starting web server on port %s\n", port)
	http.ListenAndServe(port, nil)
}
