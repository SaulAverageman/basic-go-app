package handler

import (
	"net/http"

	"github.com/saulaverageman/basic-go-app/pkg/render"
)

// render home tmpl
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.tmpl")
}

// renders about tmpl
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.tmpl")
}
