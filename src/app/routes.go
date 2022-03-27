package app

import (
	"github.com/gorilla/mux"
)

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) *AppBase {
	router.HandleFunc("/", app.Home).Methods("GET")
	router.HandleFunc("/auth/{provider}", app.Auth).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", app.AuthCallback).Methods("GET")
	return app
}