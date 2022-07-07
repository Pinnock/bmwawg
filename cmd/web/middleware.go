package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   cfg.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoadAndSave(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
