package controllers

import (
	. "github.com/giftxtrade/api/src/postgres/public/table"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func (ctr Controller) TestRoute(c *fiber.Ctx) error {
	query := SELECT(
		Event.AllColumns,
		String(slug.Make("blah")).AS("event.slug"),
		Link.AllColumns,
		User.AllColumns,
		Participant.AllColumns,
	).FROM(
		Event.
			LEFT_JOIN(Link, Event.ID.EQ(Link.EventID)).
			LEFT_JOIN(Participant, Event.ID.EQ(Participant.EventID)).
			LEFT_JOIN(User, Participant.UserID.EQ(User.ID)),
	).ORDER_BY(
		Event.DrawAt.ASC(),
		Event.CloseAt.ASC(),
		Participant.ID.ASC(),
	)

	var dest []types.Event
	err := query.QueryContext(c.Context(), ctr.DB, &dest)
	if err != nil {
		return utils.FailResponse(c, err.Error())
	}
	return utils.DataResponse(c, dest)
}
