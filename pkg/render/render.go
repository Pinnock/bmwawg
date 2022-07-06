package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/pinnock/bmwawg/pkg/config"
	"github.com/pinnock/bmwawg/pkg/models"
)

var appConf *config.AppConfig

func SetConfig(cfg *config.AppConfig) {
	appConf = cfg
}

func addDefaultData(td *models.TemplateData) {

}

func RenderTemplate(
	w http.ResponseWriter,
	tmpl string,
	d *models.TemplateData,
) {
	var cache map[string]*template.Template
	var err error

	if appConf.UseCache {
		cache = appConf.TemplateCache
	} else {
		cache, err = CreateTemplateCache()
		if err != nil {
			log.Fatalln("failed to create template cache")
		}
	}

	t, ok := cache[tmpl]
	if !ok {
		log.Fatalf("template %s not found in template cache\n", tmpl)
	}

	addDefaultData(d)

	if err := t.Execute(w, d); err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return cache, err
	}

	for _, pg := range pages {
		name := filepath.Base(pg)
		ts, err := template.New(name).ParseFiles(pg)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}
	return cache, nil
}
