package types

type Response struct {
	Message string `json:"message"`
}

type Result struct {
	Data interface{} `json:"data"`
}

type Errors struct {
	Errors interface{} `json:"errors"`
}

type DbConnection struct {
	DbName string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type TwitterKeys struct {
	ApiKey string `json:"api_key"`
	ApiKeySecret string `json:"api_key_secret"`
	BearerToken string `json:"bearer_token"`
}

type GoogleKeys struct {
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type SendgridKeys struct {
	ApiKey string `json:"api_key"`
}

type Tokens struct {
	JwtKey string `json:"jwt_key"`
	Twitter TwitterKeys `json:"twitter"`
	Google GoogleKeys `json:"google"`
	Sendgrid SendgridKeys `json:"sendgrid"`
	// To add other tokens create a struct and add them here,
	// make sure to also update tokens.json
}

type Auth struct {
	User User `json:"user"`
	Token string `json:"token"`
}

type CreateUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ImageUrl string `json:"image_url"`
}