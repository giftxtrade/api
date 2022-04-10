package app

import (
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
)

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) *AppBase {
	router.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, types.Response{
			Message: "GiftTrade REST API ⚡",
		})
	})).Methods("GET")

	// auth routes
	router.Handle("/auth/profile", app.UseJwtAuth(http.HandlerFunc(app.GetProfile))).Methods("GET")
	router.HandleFunc("/auth/{provider}", app.Auth).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", app.AuthCallback).Methods("GET")

	// products routes
	router.Handle("/products", app.UseAdminOnly(http.HandlerFunc(app.CreateProduct))).Methods("POST")

	return app
}