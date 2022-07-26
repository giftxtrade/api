package types

import "time"

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

type Auth struct {
	User User `json:"user"`
	Token string `json:"token"`
}

type CreateUser struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	ImageUrl string `json:"imageUrl" validate:"omitempty,url"`
}

type CreateCategory struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Url string `json:"url" validate:"omitempty,url"`
}

type CreateProduct struct {
	Title string `json:"title" validate:"required"`
	Description string `json:"description"`
	ProductKey string `json:"productKey" validate:"required"`
	ImageUrl string `json:"imageUrl" validate:"omitempty,url"`
	Rating float32 `json:"rating" validate:"required,min=1,max=5"`
	Price float32 `json:"price" validate:"required,gte=1"`
	OriginalUrl string `json:"originalUrl" validate:"required,url"`
	TotalReviews uint `json:"totalReviews" validate:"required,gte=1"`
	Category string `json:"category" validate:"required"`
}

type ProductFilter struct {
	Search string `json:"search" validate:"omitempty"`
	Limit int `json:"limit" validate:"required,min=1,max=200"`
	Page int `json:"page" validate:"required,gte=1"`
	MinPrice float32 `json:"minPrice" validate:"omitempty,gte=1,ltefield=MaxPrice"`
	MaxPrice float32 `json:"maxPrice" validate:"omitempty,gtefield=MinPrice"`
	Sort string `json:"sort" validate:"omitempty"`
}

type CreateEvent struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description"`
	Budget float32 `json:"budget" validate:"required,gte=1"`
	InviteMessage string `json:"inviteMessage"`
	DrawAt time.Time `json:"drawAt" validate:"required,datetime"`
	CloseAt time.Time `json:"closeAt" validate:"required,datetime"`
	CreatedBy User `json:"created_by" validate:"required"`
}