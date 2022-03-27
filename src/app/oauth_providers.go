package app

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/twitter"
)

func (app *AppBase) SetupOauthProviders() {
	goth.UseProviders(twitter.New(app.Tokens.Twitter.ApiKey, app.Tokens.Twitter.ApiKeySecret, "http://localhost:3001/auth/twitter/callback"))
}