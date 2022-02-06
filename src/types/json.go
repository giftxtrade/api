package types

type Response struct {
	Message string `json:"message"`
}

type Result struct {
	Data interface{} `json:"data"`
}

type DbConnection struct {
	DbName string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type CreatePost struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

type CreateUser struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TwitterKeys struct {
	ApiKey string `json:"api_key"`
	ApiKeySecret string `json:"api_key_secret"`
	BearerToken string `json:"bearer_token"`
}

type Tokens struct {
	JwtKey string `json:"jwt_key"`
	Twitter TwitterKeys `json:"twitter"`
	// To add other tokens create a struct and add them here,
	// make sure to also update tokens.json
}

type Auth struct {
	User User `json:"user"`
	Token string `json:"token"`
}
