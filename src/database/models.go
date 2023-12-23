// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type CurrencyType string

const (
	CurrencyTypeUSD CurrencyType = "USD"
	CurrencyTypeCAD CurrencyType = "CAD"
)

func (e *CurrencyType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CurrencyType(s)
	case string:
		*e = CurrencyType(s)
	default:
		return fmt.Errorf("unsupported scan type for CurrencyType: %T", src)
	}
	return nil
}

type NullCurrencyType struct {
	CurrencyType CurrencyType `json:"currencyType"`
	Valid        bool         `json:"valid"` // Valid is true if CurrencyType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrencyType) Scan(value interface{}) error {
	if value == nil {
		ns.CurrencyType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CurrencyType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrencyType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CurrencyType), nil
}

func (e CurrencyType) Valid() bool {
	switch e {
	case CurrencyTypeUSD,
		CurrencyTypeCAD:
		return true
	}
	return false
}

func AllCurrencyTypeValues() []CurrencyType {
	return []CurrencyType{
		CurrencyTypeUSD,
		CurrencyTypeCAD,
	}
}

type Category struct {
	ID          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
	CategoryUrl sql.NullString `db:"category_url" json:"categoryUrl"`
	CreatedAt   time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updatedAt"`
}

type Draw struct {
	ID        int64     `db:"id" json:"id"`
	DrawerID  int64     `db:"drawer_id" json:"drawerId"`
	DraweeID  int64     `db:"drawee_id" json:"draweeId"`
	EventID   int64     `db:"event_id" json:"eventId"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

type Event struct {
	ID                int64          `db:"id" json:"id"`
	Name              string         `db:"name" json:"name"`
	Description       sql.NullString `db:"description" json:"description"`
	Budget            string         `db:"budget" json:"budget"`
	InvitationMessage string         `db:"invitation_message" json:"invitationMessage"`
	DrawAt            time.Time      `db:"draw_at" json:"drawAt"`
	CloseAt           time.Time      `db:"close_at" json:"closeAt"`
	CreatedAt         time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updatedAt"`
}

type Link struct {
	ID             int64     `db:"id" json:"id"`
	Code           string    `db:"code" json:"code"`
	ExpirationDate time.Time `db:"expiration_date" json:"expirationDate"`
	EventID        int64     `db:"event_id" json:"eventId"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`
}

type Participant struct {
	ID           int64          `db:"id" json:"id"`
	Name         string         `db:"name" json:"name"`
	Email        string         `db:"email" json:"email"`
	Address      sql.NullString `db:"address" json:"address"`
	Organizer    bool           `db:"organizer" json:"organizer"`
	Participates bool           `db:"participates" json:"participates"`
	Accepted     bool           `db:"accepted" json:"accepted"`
	EventID      int64          `db:"event_id" json:"eventId"`
	UserID       sql.NullInt64  `db:"user_id" json:"userId"`
	CreatedAt    time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time      `db:"updated_at" json:"updatedAt"`
}

type Product struct {
	ID           int64          `db:"id" json:"id"`
	Title        string         `db:"title" json:"title"`
	Description  sql.NullString `db:"description" json:"description"`
	ProductKey   string         `db:"product_key" json:"productKey"`
	ImageUrl     string         `db:"image_url" json:"imageUrl"`
	TotalReviews int32          `db:"total_reviews" json:"totalReviews"`
	Rating       float32        `db:"rating" json:"rating"`
	Price        string         `db:"price" json:"price"`
	Currency     CurrencyType   `db:"currency" json:"currency"`
	Url          string         `db:"url" json:"url"`
	CategoryID   sql.NullInt64  `db:"category_id" json:"categoryId"`
	CreatedAt    time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time      `db:"updated_at" json:"updatedAt"`
	ProductTs    interface{}    `db:"product_ts" json:"productTs"`
	Origin       string         `db:"origin" json:"origin"`
}

type User struct {
	ID        int64          `db:"id" json:"id"`
	Name      string         `db:"name" json:"name"`
	Email     string         `db:"email" json:"email"`
	ImageUrl  string         `db:"image_url" json:"imageUrl"`
	Phone     sql.NullString `db:"phone" json:"phone"`
	Admin     bool           `db:"admin" json:"admin"`
	Active    bool           `db:"active" json:"active"`
	CreatedAt time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time      `db:"updated_at" json:"updatedAt"`
}

type Wish struct {
	ID            int64         `db:"id" json:"id"`
	UserID        int64         `db:"user_id" json:"userId"`
	ParticipantID int64         `db:"participant_id" json:"participantId"`
	ProductID     sql.NullInt64 `db:"product_id" json:"productId"`
	EventID       int64         `db:"event_id" json:"eventId"`
	CreatedAt     time.Time     `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time     `db:"updated_at" json:"updatedAt"`
}
