package main

import (
	"log"
	"net/http"

	"github.com/saulaverageman/basic-go-app/pkg/config"
	"github.com/saulaverageman/basic-go-app/pkg/handler"
	"github.com/saulaverageman/basic-go-app/pkg/render"
)

func main() {
	// app is the app config for the whole project
	var app config.AppConfig

	// render template cache as soon as the app starts and saving them in appConfig.TemplateCache
	tc, err := render.FormTemplateCache("/work/")
	if err != nil {
		log.Fatal("error: cannot create template cache")
	} else {
		app.TemplateCache = tc
	}

	// sending app config to render.go
	render.SetRenderAppConfig(&app)

	const appPort = ":8000"
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	_ = http.ListenAndServe(appPort, nil)

}
