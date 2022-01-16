package routes

import (
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
)

func (app *AppBase) Home(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	if message == "" {
		utils.JsonResponse(w, types.Response{
			Message: "Hello world!",
		})
		return
	}
	utils.JsonResponse(w, types.Response{
		Message: message,
	})
}