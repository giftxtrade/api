package controllers

import (
	"net/http"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type AuthController struct {
	Controller
	UserServices *services.UserService
}

func (ctx *AuthController) CreateRoutes(router *mux.Router, path string) {
	router.Handle(path + "/profile", utils.UseJwtAuth(ctx.Tokens.JwtKey, ctx.UserServices, http.HandlerFunc(ctx.get_profile))).Methods("GET")
	router.HandleFunc(path + "/{provider}", ctx.sign_in).Methods("GET")
	router.HandleFunc(path + "/{provider}/callback", ctx.callback).Methods("GET")
}

func (ctx *AuthController) get_profile(w http.ResponseWriter, r *http.Request) {
	auth := utils.ParseAuthContext(r.Context())
	utils.DataResponse(w, &auth)
}


// [GET] /auth/{provider}
func (ctx *AuthController) sign_in(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	provider := params["provider"]
	
	callback_url := r.URL.Query().Get("callbackUrl")
	if callback_url != "" {
		switch provider {
		case "google":
			goth.UseProviders(utils.CreateGoogleProvider(callback_url, ctx.Tokens.Google))
		}
	}

	gothic.BeginAuthHandler(w, r)
}

// [GET] /auth/{provider}/callback
func (ctx *AuthController) callback(w http.ResponseWriter, r *http.Request) {
	provider_user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		utils.FailResponse(w, "could not complete authentication")
		return
	}

	check_user := types.CreateUser{
		Email: provider_user.Email,
		Name: provider_user.Name,
		ImageUrl: provider_user.AvatarURL,
	}
	user := ctx.UserServices.FindOrCreate(&check_user)
	token, err := utils.GenerateJWT(ctx.Tokens.JwtKey, user)
	if err != nil {
		utils.FailResponse(w, "could not generate token")
		return
	}
	auth := types.Auth{
		Token: token,
		User: *user,
	}
	utils.JsonResponse(w, &auth)
}