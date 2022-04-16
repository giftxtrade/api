package controllers

import "github.com/giftxtrade/api/src/types"

type Controller struct {
	types.AppContext
}

type IController interface {
	CreateController()
}