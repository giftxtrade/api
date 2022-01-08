package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ayaanqui/go-rest-server/src/types"
	"github.com/ayaanqui/go-rest-server/src/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// [POST] /register controller
func (app *AppBase) Register(w http.ResponseWriter, r *http.Request) {
	post_user := types.CreateUser{}
	if err := json.NewDecoder(r.Body).Decode(&post_user); err != nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: "Incorrect json formatting"})
		return
	}

	// Verify if username or email already exists
	check := types.User{}
	app.DB.Table("users").Find(
		&check, 
		"username = ? OR email = ?", 
		post_user.Username, 
		post_user.Email,
	)
	if check.ID != uuid.Nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: "An account with the username or email already exists"})
		return
	}

	// User does not exist, so create the account
	new_user := types.User{
		Username: post_user.Username,
		Email: post_user.Email,
		Password: post_user.Password,
		IsAdmin: false,
		IsActive: true,
	}
	app.DB.Table("users").Create(&new_user)
	utils.JsonResponse(w, &new_user)
}

// [POST] /login controller
func (app *AppBase) Login(w http.ResponseWriter, r *http.Request) {
	login := types.LoginUser{}
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: "Could not parse json"})
		return
	}

	// Fetch user with username field
	const message string = "Username or password is incorrect"
	user := types.User{}
	app.DB.Table("users").Find(&user, "username = ? OR email = ?", login.Username, login.Username)
	if user.ID == uuid.Nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: message})
		return
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: message})
		return
	}
	
	// Generate JWT token
	token, err := generate_token([]byte(app.Tokens.JwtKey), &user)
	if err != nil {
		w.WriteHeader(400)
		utils.JsonResponse(w, types.Response{Message: "Could not generate token"})
		return
	}
	utils.JsonResponse(w, types.Auth{
		Token: token,
		User: user,
	})
}

func generate_token(key []byte, user *types.User) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"username": user.Username,
		"email": user.Email,
		"is_active": user.IsActive,
	})
	token, err := jwt.SignedString(key)
	if err != nil {
		return "", err
	}
	return token, nil
}

// [GET] /me controller
func (app *AppBase) Profile(w http.ResponseWriter, r *http.Request) {
	auth := r.Context().Value(types.AuthKey).(types.Auth)
	utils.JsonResponse(w, auth.User)
}