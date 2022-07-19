package controllers

import (
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/gorilla/mux"
)

type Controller struct {
	types.AppContext
	Service services.Service
}

type IController interface {
	CreateController(router *mux.Router, path string)
}

func New(app_ctx types.AppContext, service services.Service) Controller {
	controller := Controller{
		AppContext: app_ctx,
		Service: service,
	}
	server := app_ctx.Server

	// create routes
	server.Get("/", controller.Home)
	auth := server.Group("/auth")
	{ // auth
		auth.Get("/profile", controller.UseJwtAuth, controller.GetProfile)
		auth.Get("/:provider", controller.SignIn)
		auth.Get("/:provider/callback", controller.Callback)
	}

	products := server.Group("/products")
	{
		products.Post("", controller.UseAdminOnly, controller.CreateProduct)
		products.Get("", controller.UseAdminOnly, controller.FindAllProducts)
		products.Get("/:id", controller.UseJwtAuth, controller.FindProduct)
	}
	server.Get("*", controller.NotFound)
	return controller
}