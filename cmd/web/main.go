package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pinnock/bmwawg/internal/config"
	"github.com/pinnock/bmwawg/internal/render"
)

const port string = ":8080"

var cfg *config.AppConfig
var session *scs.SessionManager

func main() {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("failed to create template cache: %v\n", err)
	}

	cfg = &config.AppConfig{
		TemplateCache: tc,
		UseCache:      false,
	}

	render.SetConfig(cfg)
	session = scs.New()
	session.Lifetime = 8 * time.Hour
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = cfg.InProduction

	cfg.Session = session

	svr := &http.Server{
		Addr:    port,
		Handler: routes(cfg),
	}

	log.Printf("starting web server on port %s\n", port)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
