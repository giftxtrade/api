//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type EventLink struct {
	ID                 *int64
	Name               *string
	Description        *string
	Budget             *string
	InvitationMessage  *string
	DrawAt             *time.Time
	CloseAt            *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	LinkID             *int64
	LinkCode           *string
	LinkExpirationDate *time.Time
}