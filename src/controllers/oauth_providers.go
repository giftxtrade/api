package controllers

import (
	"github.com/giftxtrade/api/src/types"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitter"
)

func SetupOauthProviders(tokens types.Tokens) {
	goth.UseProviders(
		CreateTwitterProvider(tokens.Twitter.CallbackUrl, tokens.Twitter), 
		CreateGoogleProvider(tokens.Google.CallbackUrl, tokens.Google),
	)
}

func CreateTwitterProvider(callback_url string, tokens types.TwitterKeys) *twitter.Provider {
	return twitter.New(tokens.ApiKey, tokens.ApiKeySecret, callback_url)
}

func CreateGoogleProvider(callback_url string, tokens types.GoogleKeys) *google.Provider {
	return google.New(tokens.ClientId, tokens.ClientSecret, callback_url, "profile", "email")
}
