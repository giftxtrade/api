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
	DbName string `json:"dbName"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type TwitterKeys struct {
	ApiKey string `json:"apiKey"`
	ApiKeySecret string `json:"apiKeySecret"`
	BearerToken string `json:"bearerToken"`
}

type GoogleKeys struct {
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
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

type Auth struct {
	User User `json:"user"`
	Token string `json:"token"`
}

type CreateUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ImageUrl string `json:"imageUrl"`
}

type CreateCategory struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Url string `json:"url"`
}

type CreateProduct struct {
	Title string `json:"title"`
	Description string `json:"description"`
	ProductKey string `json:"productKey"`
	ImageUrl string `json:"imageUrl"`
	Rating float32 `json:"rating"`
	Price float32 `json:"price"`
	OriginalUrl string `json:"originalUrl"`
	WebsiteOrigin string `json:"websiteOrigin"`
	TotalReviews int `json:"totalReviews"`
	Category string `json:"category"`
}