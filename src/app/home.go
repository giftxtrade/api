package app

import (
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
)

func (app *AppBase) Home(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, types.Response{
		Message: "GiftTrade API ⚡",
	})
}