package render

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles(
		"./templates/"+tmpl,
		"./templates/base.layout.gohtml",
	)
	if err != nil {
		log.Printf("Error parsing template files: %v\n", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("Error applying parsed template: %v\n", err)
	}
}
