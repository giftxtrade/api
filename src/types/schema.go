package types

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Base struct {
	ID uuid.UUID `gorm:"type:uuid; primary key" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New()
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return nil
}

func (base *Base) BeforeUpdate(tx *gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil
}

type Post struct {
	Base
	Title string `gorm:"not null" json:"title"`
	Slug string `gorm:"not null" json:"slug"`
	Content string `gorm:"type:text; not null" json:"content"`
	Summary string `gorm:"not null; default: ''" json:"summary"`
}

type User struct {
	Base
	Username string `gorm:"varchar(30); not null; index; unique" json:"username"`
	Email string `gorm:"varchar(255); not null; index; unique" json:"email"`
	// Ignore password field from json output
	Password string `gorm:"varchar(255); index" json:"-"`
	IsAdmin bool `gorm:"default: false" json:"is_admin"`
	IsActive bool `gorm:"default: false" json:"is_active"`
}

// Hashes password using the bcrypt library
func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Base.BeforeCreate(tx) // Call base BeforeCreate first
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}