package controllers

import (
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
)

type Controller struct {
	types.AppContext
	Service services.Service
	Querier *database.Queries
}

type IController interface {
	New(app_ctx types.AppContext, service services.Service) Controller
}

func New(app_ctx types.AppContext, querier *database.Queries, service services.Service) Controller {
	controller := Controller{
		AppContext: app_ctx,
		Service: service,
		Querier: querier,
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
		products.Get("", controller.UseJwtAuth, controller.FindAllProducts)
		products.Get("/:id", controller.UseJwtAuth, controller.FindProduct)
	}

	events := server.Group("/events")
	{
		events.Post("", controller.UseJwtAuth, controller.CreateEvent)
		events.Get("", controller.UseJwtAuth, controller.GetAllEvents)
	}

	server.Get("*", controller.NotFound)
	return controller
}