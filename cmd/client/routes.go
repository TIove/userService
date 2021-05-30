package main

import (
	"net/http"
)

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	userRoutes(app, mux)

	return mux
}

func userRoutes(app *Application, mux *http.ServeMux) {
	mux.HandleFunc("/user/create", app.create)
	mux.HandleFunc("/user/delete", app.delete)
	mux.HandleFunc("/user/get", app.get)
}