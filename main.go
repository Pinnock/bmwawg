package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

// Home is the handler for the Home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page\n")
}

// About is the handler for the About page
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About page\n")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32
	x = 100
	y = 0
	q, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	fmt.Fprintf(w, "%f divided by %f = %f\n", x, y, q)
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return (x / y), nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	log.Printf("Starting web server on port %s\n", port)
	http.ListenAndServe(port, nil)
}
