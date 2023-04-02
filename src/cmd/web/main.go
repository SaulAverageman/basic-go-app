package main

import (
	"log"
	"net/http"

	"github.com/saulaverageman/basic-go-app/pkg/config"
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
	render.NewRender(&app)

	const appPort = ":8000"

	server := &http.Server{
		Addr:    appPort,
		Handler: routes(&app),
	}

	log.Println("verbose: starting servet @port:", appPort)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
