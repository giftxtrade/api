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
	{
		auth := server.Group("/auth")
		profile := auth.Group("/profile")
		{
			profile.Use(controller.UseJwtAuth)
			profile.Get("", controller.GetProfile)
		}
		auth.Get("/:provider", controller.SignIn)
		auth.Get("/:provider/callback", controller.Callback)
	}
	return controller
}