package routes

import (
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
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
	utils.JsonResponse(w, user_from_db)
}

// AUTH REQUIRED - [GET] /auth/profile
func (app *AppBase) Profile(w http.ResponseWriter, r *http.Request) {
	auth := r.Context().Value(types.AuthKey).(types.Auth)
	utils.JsonResponse(w, auth.User)
}