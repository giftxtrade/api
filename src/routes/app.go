package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/twitter"
	"gorm.io/gorm"
)

type AppBase struct {
	DB *gorm.DB
	Tokens types.Tokens
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
	CreateRoutes(router *mux.Router)
}

func (app *AppBase) NewBaseHandler(conn *gorm.DB) *AppBase {
	app.DB = conn
	tokens, tokens_err := utils.ParseTokens()
	if tokens_err != nil {
		panic(tokens_err)
	}
	app.Tokens = tokens

	err := conn.AutoMigrate(
		&types.User{},
	)
	if err != nil {
		log.Fatal("Could not generate schema.\n")
		panic(tokens_err)
	}
	return app
}

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) *AppBase {
	// Initialize goth
	goth.UseProviders(twitter.New(app.Tokens.Twitter.ApiKey, app.Tokens.Twitter.ApiKeySecret, "http://localhost:3001/auth/twitter/callback"))

	router.HandleFunc("/", app.Home).Methods("GET")
	
	// Auth routes
	router.Handle("/auth/profile", UseJwtAuth(app, http.HandlerFunc(app.Profile))).Methods("GET")
	router.HandleFunc("/auth/{provider}", app.Auth).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", app.AuthCallback).Methods("GET")

	router.Handle("/admin", AdminOnly(app, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, types.Response{Message: "Admin only page"})
	}))).Methods("GET")
	return app
}

func New(conn *gorm.DB) *AppBase {
	app := AppBase{}
	return app.NewBaseHandler(conn)
}