package app

import (
	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/types"
	"github.com/gorilla/mux"
)

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) *AppBase {
	controller := controllers.Controller{
		AppContext: types.AppContext{
			DB: app.DB,
			Tokens: app.Tokens,
		},
	}

	home_controller := controllers.HomeController{
		Controller: controller,
	}
	home_controller.CreateRoutes(router, "/")

	auth_controller := controllers.AuthController{
		Controller: controller,
		UserServices: app.UserServices,
	}
	auth_controller.CreateRoutes(router, "/auth")

	products_controller := controllers.ProductsController{
		Controller: controller,
		UserServices: app.UserServices,
		ProductServices: app.ProductServices,
	}
	products_controller.CreateRoutes(router, "/products")

	return app
}