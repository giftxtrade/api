package routes

import (
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
	"github.com/markbates/goth/gothic"
)

// [GET] /auth/{provider}
func (app *AppBase) Auth(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

// [GET] /auth/{provider}/callback
func (app *AppBase) AuthCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: "Could not complete authentication"})
		return
	}
	utils.JsonResponse(w, types.Response{Message: "Hello, " + user.Name})
}