package mappers

import (
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

func DbLinkToLink(link database.Link, event *database.Event) types.Link {
	mapped_link := types.Link{
		ID: link.ID,
		Code: link.Code,
		ExpirationDate: link.ExpirationDate,
		EventID: link.EventID,
	}
	if event != nil {
		event := DbEventToEvent(*event, nil, nil)
		mapped_link.Event = &event
	}
	return mapped_link
} 
