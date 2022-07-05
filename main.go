package main

import (
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
