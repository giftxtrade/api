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
	}
	if event != nil {
		mapped_link.EventID = event.ID
	}
	return mapped_link
} 
