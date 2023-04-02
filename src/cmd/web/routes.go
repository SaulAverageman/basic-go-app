package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/saulaverageman/basic-go-app/pkg/config"
	"github.com/saulaverageman/basic-go-app/pkg/handler"
)

var appConfig *config.AppConfig

func routes(app *config.AppConfig) http.Handler {
	appConfig = app
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handler.Home))
	// mux.Get("/about", http.HandlerFunc(handler.About))

	mux := chi.NewMux()
	mux.Get("/", handler.Home)
	mux.Get("/about", handler.About)

	return mux
}
