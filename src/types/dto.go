package types

import (
	"time"
)

type Response struct {
	Message string `json:"message"`
}

type Result struct {
	Data interface{} `json:"data"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

type DeleteStatus struct {
	Deleted bool `json:"deleted"`
}

type User struct {
	ID int64 `json:"id" sql:"primary_key"`
	Name string `json:"name"`
	Email string `json:"email" `
	ImageUrl string `json:"imageUrl,omitempty"`
	Active bool `json:"active"`
	Phone string `json:"phone,omitempty"`
	Admin bool `json:"admin,omitempty"`
}

type CreateUser struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	ImageUrl string `json:"imageUrl,omitempty" validate:"omitempty,url"`
	Phone string `json:"phone,omitempty" validate:"omitempty,"`
}

type Auth struct {
	User User `json:"user"`
	Token string `json:"token"`
}

type Category struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
}

type CreateCategory struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty" validate:"omitempty,url"`
}

type Product struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
	ProductKey string `json:"productKey"`
	ImageUrl string `json:"imageUrl"`
	TotalReviews int32 `json:"totalReviews"`
	Rating float32 `json:"rating"`
	Price string `json:"price"`
	Currency string `json:"currency"`
	Url string `json:"url"`
	CategoryID int64  `json:"categoryId,omitempty"`
	Category Category `json:"category,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Origin string `json:"origin"`
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
	Search *string `json:"search,omitempty" validate:"omitempty"`
	Limit int32 `json:"limit" validate:"required,min=1,max=200"`
	Page int32 `json:"page" validate:"required,gte=1"`
	MinPrice float32 `json:"minPrice,omitempty" validate:"omitempty,gte=1,ltefield=MaxPrice"`
	MaxPrice float32 `json:"maxPrice,omitempty" validate:"omitempty,gtefield=MinPrice"`
	Sort *string `json:"sort,omitempty" validate:"omitempty"`
}

type CreateWish struct {
	ProductID *int64 `json:"productId,omitempty"`
}

type DeleteWish struct {
	WishID int64 `json:"wishId"`
}

type Wish struct {
	ID int64 `json:"id" sql:"primary_key"`
	UserID int64 `json:"userId"`
	ParticipantID int64 `json:"participantId"`
	ProductID int64 `json:"productId,omitempty"`
	Product *Product `json:"product,omitempty"`
	EventID int64 `json:"eventId"`
	Quantity int32 `json:"quantity" alias:"wish.quantity"`
}

type Participant struct {
	ID int64 `json:"id" sql:"primary_key"`
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address,omitempty"`
	Organizer bool `json:"organizer"`
	Participates bool `json:"participates"`
	Accepted bool `json:"accepted"`
	EventID int64 `json:"eventId,omitempty"`
	Event *Event `json:"event,omitempty"`
	UserID int64 `json:"userId,omitempty"`
	User *User `json:"user,omitempty"`
	Wishes *[]Wish `json:"wishes,omitempty"`
}

type CreateParticipant struct {
	Email string `json:"email" validate:"required,email"`
	Name string `json:"name,omitempty" validate:"omitempty"`
	Address string `json:"address,omitempty" validate:"omitempty"`
	Organizer bool `json:"organizer,omitempty" validate:"omitempty"`
	Participates bool `json:"participates,omitempty" validate:"omitempty"`
}

type PatchParticipant struct {
	Name *string `json:"name,omitempty" validate:"omitempty"`
	Address *string `json:"address,omitempty" validate:"omitempty"`
	Organizer *bool `json:"organizer,omitempty" validate:"omitempty"`
	Participates *bool `json:"participates,omitempty" validate:"omitempty"`
}

type Link struct {
	ID int64 `json:"id" sql:"primary_key" alias:"link.id"`
	Code string `json:"code" alias:"link.code"`
	EventID int64 `json:"eventId,omitempty" alias:"link.event_id"`
	Event *Event `json:"event,omitempty"`
	ExpirationDate time.Time `json:"expiration_date" alias:"link.expiration_date"`
}

type Event struct {
	ID int64 `json:"id" sql:"primary_key" alias:"event.id"`
	Name string `json:"name" alias:"event.name"`
	Slug string `json:"slug,omitempty"`
	Description string `json:"description,omitempty" alias:"event.description"`
	Budget string `json:"budget" alias:"event.budget"`
	InvitationMessage string `json:"invitationMessage,omitempty" alias:"event.invitation_message"`
	DrawAt time.Time `json:"drawAt" alias:"event.draw_at"`
	CloseAt time.Time `json:"closeAt" alias:"event.close_at"`
	CreatedAt time.Time `json:"createdAt" alias:"event.created_at"`
	UpdatedAt time.Time `json:"updatedAt" alias:"event.updated_at"`
	Participants []Participant `json:"participants,omitempty"`
	Links []Link `json:"links,omitempty"`
	MyWishList []Wish `json:"my_wish_list,omitempty"`
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

type UpdateEvent struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Budget float32 `json:"budget,omitempty" validate:"gte=1"`
	DrawAt time.Time `json:"drawAt,omitempty"`
	CloseAt time.Time `json:"closeAt,omitempty"`
}
