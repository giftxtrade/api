package app

import (
	"net/http"

	"github.com/giftxtrade/api/src/services"
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
	provider_user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		utils.FailResponse(w, "could not complete authentication")
		return
	}

	check_user := types.User{
		Email: provider_user.Email,
		Name: provider_user.Name,
		ImageUrl: provider_user.AvatarURL,
	}
	user := services.GetUserByEmailOrCreate(app.DB, &check_user)
	token, err := utils.GenerateJWT(app.Tokens.JwtKey, &user)
	if err != nil {
		utils.FailResponse(w, "could not generate token")
		return
	}
	auth := types.Auth{
		Token: token,
		User: user,
	}
	utils.JsonResponse(w, auth)
}

// Auth required [GET] /auth/profile
func (app *AppBase) GetProfile(w http.ResponseWriter, r *http.Request) {
	auth := utils.ParseAuthContext(r.Context())
	utils.DataResponse(w, &auth)
}