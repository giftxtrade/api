package app

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitter"
)

func (app *AppBase) SetupOauthProviders() {
	goth.UseProviders(
		app.CreateTwitterProvider(""), 
		app.CreateGoogleProvider("https://giftxtrade.com/auth/google/callback"),
	)
}

func (app *AppBase) CreateTwitterProvider(callback_url string) *twitter.Provider {
	tokens := app.Tokens.Twitter
	return twitter.New(tokens.ApiKey, tokens.ApiKeySecret, callback_url)
}

func (app *AppBase) CreateGoogleProvider(callback_url string) *google.Provider {
	tokens := app.Tokens.Google
	return google.New(tokens.ClientId, tokens.ClientSecret, callback_url, "profile")
}