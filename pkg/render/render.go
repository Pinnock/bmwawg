package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate2(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles(
		"./templates/"+tmpl,
		"./templates/base.layout.gohtml",
	)
	if err != nil {
		log.Printf("error parsing template files: %v\n", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("error applying parsed template: %v\n", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	// check to see if template is already isCached
	_, isCached := tc[t]
	if !isCached {
		log.Printf("creating and caching template: %s\n", t)
		if err := cacheTemplate(t); err != nil {
			log.Println(err)
		}
	} else {
		log.Printf("using cached template: %s\n", t)
	}

	if err := tc[t].Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func cacheTemplate(t string) error {
	tt := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gohtml",
	}

	tmpl, err := template.ParseFiles(tt...)
	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}
