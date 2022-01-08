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
	AccessToken string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

type Tokens struct {
	JwtKey string `json:"jwt_key"`
	Twitter TwitterKeys `json:"twitter"`
}

type Auth struct {
	User User `json:"user"`
	Token string `json:"token"`
}