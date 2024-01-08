package controllers

import (
	"context"
	"fmt"
	"strconv"

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
		events.Get("/verify-invite-code/:invite_code", c.VerifyEventLinkCode)
		events.Get("/join/:invite_code", c.UseJwtAuth, c.JoinEventViaInviteCode)
		events.Post("", c.UseJwtAuth, c.CreateEvent)
		events.Get("", c.UseJwtAuth, c.GetEvents)
		events.Get("/invites", c.UseJwtAuth, c.GetInvites)
		events.Get("/invites/accept/:event_id", c.UseJwtAuth, c.UseEventAuthWithParam, c.AcceptEventInvite)
		events.Get("/invites/decline/:event_id", c.UseJwtAuth, c.UseEventAuthWithParam, c.DeclineEventInvite)
		events.Get("/get-link/:event_id", c.UseJwtAuth, c.UseEventAuthWithParam, c.GetEventLink)
		events.Get("/:event_id", c.UseJwtAuth, c.UseEventAuthWithParam, c.GetEventById)
		events.Patch("/:event_id", c.UseJwtAuth, c.UseEventOrganizerAuthWithParam, c.UpdateProduct)
		events.Delete("/:event_id", c.UseJwtAuth, c.UseEventOrganizerAuthWithParam, c.DeleteEvent)
	}
	participants := server.Group("/participants")
	{
		participants.Patch("/manage/:event_id", c.UseJwtAuth, c.UseEventOrganizerAuthWithParam, c.UseEventParticipantAuthWithQuery, c.ManageParticipantUpdate)
		participants.Delete("/manage/:event_id", c.UseJwtAuth, c.UseEventOrganizerAuthWithParam, c.UseEventParticipantAuthWithQuery, c.ManageParticipantRemoval)
		participants.Get("/:event_id/:participant_id", c.UseJwtAuth, c.UseEventAuthWithParam, c.UseEventParticipantAuthWithParam, c.GetParticipantById)
	}
	server.Get("*", func(c *fiber.Ctx) error {
		return utils.ResponseWithStatusCode(c, fiber.ErrNotFound.Code, types.Errors{
			Errors: []string{"resource not found"},
		})
	})
	return c
}

func SetUserContext(c *fiber.Ctx, key interface{}, value interface{}) {
	c.SetUserContext(context.WithValue(c.UserContext(), key, value))
}

// Given a `fiber.Ctx.UserContext`, find and return the auth struct using the types.AuthKey key
func GetAuthContext(user_context context.Context) types.Auth {
	auth := user_context.Value(AUTH_KEY).(types.Auth)
	return auth
}

// Returns the even_id based on the route `*/:event_id/*` param
func ParseEventIdFromRoute(c *fiber.Ctx) (event_id int64, error error) {
	id_raw := c.Params("event_id")
	id, err := strconv.ParseInt(id_raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid event id")
	}
	return id, nil
}

func GetEventIdFromContext(user_context context.Context) int64 {
	id := user_context.Value(EVENT_ID_PARAM_KEY).(int64)
	return id
}
