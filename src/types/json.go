package types

type TwitterKeys struct {
	ApiKey string `json:"apiKey"`
	ApiKeySecret string `json:"apiKeySecret"`
	BearerToken string `json:"bearerToken"`
	CallbackUrl string `json:"callbackUrl"`
}

type GoogleKeys struct {
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	CallbackUrl string `json:"callbackUrl"`
}

type SendgridKeys struct {
	ApiKey string `json:"apiKey"`
}

type Tokens struct {
	JwtKey string `json:"jwtKey"`
	Twitter TwitterKeys `json:"twitter"`
	Google GoogleKeys `json:"google"`
	Sendgrid SendgridKeys `json:"sendgrid"`
	// To add other tokens create a struct and add them here,
	// make sure to also update tokens.json
}
