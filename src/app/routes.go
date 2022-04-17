package app

import (
	"net/http"

	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
)

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) *AppBase {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, types.Response{ Message: "GiftTrade REST API ⚡" })
	}).Methods("GET")

	controller := controllers.Controller{
		AppContext: types.AppContext{
			DB: app.DB,
			Tokens: app.Tokens,
		},
	}

	auth_controller := controllers.AuthController{
		Controller: controller,
		UserServices: app.UserServices,
	}
	auth_controller.CreateRoutes(router)

	products_controller := controllers.ProductsController{
		Controller: controller,
		UserServices: app.UserServices,
		ProductServices: app.ProductServices,
	}
	products_controller.CreateRoutes(router)

	return app
}