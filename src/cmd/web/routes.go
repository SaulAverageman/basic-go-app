package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/saulaverageman/basic-go-app/pkg/config"
	"github.com/saulaverageman/basic-go-app/pkg/handler"
)

var appConfig *config.AppConfig

func routes(app *config.AppConfig) http.Handler {
	appConfig = app

	mux := chi.NewMux()

	//Middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	//routes
	mux.Get("/", handler.Home)
	mux.Get("/about", handler.About)

	return mux
}
