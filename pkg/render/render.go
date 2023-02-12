package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// renders and writes templats based on given tmpl name
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/", tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error rendering template: ", tmpl, "with error:", err)
		return
	}
}
