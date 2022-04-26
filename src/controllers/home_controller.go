package controllers

import (
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
)

type HomeController struct {
	Controller
}

func (controller *HomeController) CreateRoutes(router *mux.Router, path string) {
	router.HandleFunc(path, controller.Home).Methods("GET")
}

func (controller *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, types.Response{
		Message: "GiftTrade REST API ⚡",
	})
}