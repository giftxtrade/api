package controllers

import (
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
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
	c := Controller{
		AppContext: app_ctx,
		Service: service,
		Querier: querier,
	}
	server := app_ctx.Server

	// create routes
	server.Get("/", func(c *fiber.Ctx) error {
		return utils.JsonResponse(c, types.Response{
			Message: "GiftTrade REST API ⚡",
		})
	})
	auth := server.Group("/auth")
	{
		auth.Get("/profile", c.UseJwtAuth, c.GetProfile)
		auth.Get("/google/verify", c.GoogleVerify)
		auth.Get("/:provider", c.SignIn)
		auth.Get("/:provider/callback", c.Callback)
	}
	products := server.Group("/products")
	{
		products.Post("", c.UseAdminOnly, c.CreateProduct)
		products.Get("", c.UseJwtAuth, c.FindAllProducts)
		products.Get("/:id", c.UseJwtAuth, c.FindProduct)
	}
	events := server.Group("/events")
	{
		events.Post("", c.UseJwtAuth, c.CreateEvent)
		events.Get("", c.UseJwtAuth, c.GetEvents)
		events.Get("/invites", c.UseJwtAuth, c.GetInvites)
		events.Get("/invites/accept/:event_id", c.UseJwtAuth, c.UseEventAuthWithParam, c.AcceptEventInvite)
		events.Get("/invites/decline/:event_id", c.UseJwtAuth, c.UseEventAuthWithParam, c.DeclineEventInvite)
		events.Get("/:event_id", c.UseJwtAuth, c.UseEventAuthWithParam, c.GetEventById)
	}
	server.Get("*", func(c *fiber.Ctx) error {
		return utils.ResponseWithStatusCode(c, fiber.ErrNotFound.Code, types.Errors{
			Errors: []string{"resource not found"},
		})
	})
	return c
}
