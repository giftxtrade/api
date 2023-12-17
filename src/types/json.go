package types

import "time"

type Response struct {
	Message string `json:"message"`
}

type Result struct {
	Data interface{} `json:"data"`
}

type Errors struct {
	Errors []string `json:"errors"`
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

type CreateUser struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	ImageUrl string `json:"imageUrl,omitempty" validate:"omitempty,url"`
}

type CreateCategory struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty" validate:"omitempty,url"`
}

type CreateProduct struct {
	Title string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
	ProductKey string `json:"productKey" validate:"required"`
	ImageUrl string `json:"imageUrl,omitempty" validate:"omitempty,url"`
	Rating float32 `json:"rating" validate:"required,min=1,max=5"`
	Price string `json:"price" validate:"required,gte=1"`
	OriginalUrl string `json:"originalUrl" validate:"required,url"`
	TotalReviews uint `json:"totalReviews" validate:"required,gte=1"`
	Category string `json:"category" validate:"required"`
}

type ProductFilter struct {
	Search string `json:"search,omitempty" validate:"omitempty"`
	Limit int32 `json:"limit" validate:"required,min=1,max=200"`
	Page int32 `json:"page" validate:"required,gte=1"`
	MinPrice float32 `json:"minPrice,omitempty" validate:"omitempty,gte=1,ltefield=MaxPrice"`
	MaxPrice float32 `json:"maxPrice,omitempty" validate:"omitempty,gtefield=MinPrice"`
	Sort string `json:"sort,omitempty" validate:"omitempty"`
}

type CreateEvent struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Budget float32 `json:"budget" validate:"required,gte=1"`
	InviteMessage string `json:"inviteMessage,omitempty"`
	DrawAt time.Time `json:"drawAt" validate:"required"`
	CloseAt time.Time `json:"closeAt" validate:"required"`
	Participants []CreateParticipant `json:"participants,omitempty" validate:"omitempty"`
}

type CreateParticipant struct {
	Email string `json:"email" validate:"required,email"`
	Nickname string `json:"nickname,omitempty" validate:"omitempty"`
	Address string `json:"address,omitempty" validate:"omitempty"`
	Organizer bool `json:"organizer,omitempty" validate:"omitempty"`
	Participates bool `json:"participates,omitempty" validate:"omitempty"`
}