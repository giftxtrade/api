package app

import (
	"net/http"

	"github.com/giftxtrade/api/src/controllers"
)

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes() *AppBase {
	router := app.Router
	controller := controllers.Controller{
		AppContext: app.AppContext,
		Service: app.Service,
	}

	home_controller := controllers.HomeController{
		Controller: controller,
	}
	home_controller.CreateRoutes(router, "/")

	auth_controller := controllers.AuthController{
		Controller: controller,
	}
	auth_controller.CreateRoutes(router, "/auth")

	products_controller := controllers.ProductsController{
		Controller: controller,
	}
	products_controller.CreateRoutes(router, "/products")

	// 404 page
	router.NotFoundHandler = http.HandlerFunc(home_controller.NotFound)

	return app
}