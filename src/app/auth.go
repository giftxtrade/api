package app

import (
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

// [GET] /auth/{provider}
func (app *AppBase) Auth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	provider := params["provider"]
	callback_url := r.URL.Query().Get("callback_url")

	if callback_url != "" {
		switch provider {
		case "google":
			goth.UseProviders(app.CreateGoogleProvider(callback_url))
		}
	}

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
	utils.JsonResponse(w, user)
}