package routes

import (
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
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