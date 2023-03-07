package main

import (
	"net/http"

	"github.com/saulaverageman/basic-go-app/pkg/handler"
)

func main() {
	const appPort = ":8000"
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	_ = http.ListenAndServe(appPort, nil)

}
