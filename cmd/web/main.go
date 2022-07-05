package main

import (
	"log"
	"net/http"

	"github.com/pinnock/bmwawg/pkg/handlers"
)

const port string = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting web server on port %s\n", port)
	http.ListenAndServe(port, nil)
}
