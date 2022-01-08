package types

import (
	"time"

	"github.com/google/uuid"
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
	Email string `gorm:"varchar(255); not null; index; unique" json:"email"`
	Name string `gorm:"varchar(255); not null" json:"name"`
	ImageUrl string `gorm:"varchar(255); json:"image_url"`
	IsAdmin bool `gorm:"default: false" json:"-"`
	IsActive bool `gorm:"default: false" json:"is_active"`
}