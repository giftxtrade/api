package controllers

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/gorilla/mux"
)

type Controller struct {
	types.AppContext
}

type IController interface {
	CreateController(router *mux.Router, path string)
}