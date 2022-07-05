package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!\n")
		if err != nil {
			log.Println(err)
		}
		log.Printf("Wrote %d bytes in response\n", n)
	})

	http.ListenAndServe(":8080", nil)
}
