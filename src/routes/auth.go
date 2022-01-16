package routes

import (
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

	// Handle user
	user_from_db := types.User{}
	app.DB.Table("users").Find(&user_from_db, "email = ?", user.Email)
	if user_from_db.ID == uuid.Nil {
		// user not found, so create the new user
		new_user := types.User{
			Name: user.Name,
			Email: user.Email,
			ImageUrl: user.AvatarURL,
			IsActive: true,
			IsAdmin: false,
		}
		app.DB.Create(&new_user)
		user_from_db = new_user
	}
	
	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user_from_db.ID,
		"name": user_from_db.Name,
		"email": user_from_db.Email,
		"image_url": user_from_db.ImageUrl,
	})
	signed_token, err := token.SignedString([]byte(app.Tokens.JwtKey))

	if err != nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{
			Message: "Could not authenticate",
		})
		return
	}
	// Return user and token
	utils.JsonResponse(w, types.Auth{
		User: user_from_db,
		Token: signed_token,
	})
}

// AUTH REQUIRED - [GET] /auth/profile
func (app *AppBase) Profile(w http.ResponseWriter, r *http.Request) {
	auth := r.Context().Value(types.AuthKey).(types.Auth)
	utils.JsonResponse(w, auth.User)
}