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

type User struct {
	Base
	Email string `gorm:"varchar(255); not null; index; unique" json:"email"`
	Name string `gorm:"varchar(255); not null" json:"name"`
	ImageUrl string `gorm:"varchar(255);" json:"image_url"`
	IsAdmin bool `gorm:"default: false" json:"is_admin"`
	IsActive bool `gorm:"default: false" json:"is_active"`
}

type Category struct {
	Base
	Name string `gorm:"type:varchar(30); not null; index; unique" json:"name"`
	Description string `gorm:"type:text; default: ''" json:"description"`
	Url string `gorm:"type:text" json:"url"`
}

type Product struct {
	Base
	Title string `gorm:"type:text; not null; index" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	ProductKey string `gorm:"type:varchar(255); not null; index" json:"product_key"`
	ImageUrl string `gorm:"type:text" json:"image_url"`
	Rating float32 `gorm:"type:float; not null; index" json:"rating"`
	Price float32 `gorm:"type:float(2); not null; index" json:"price"`
	OriginalUrl string `gorm:"type:text; not null" json:"original_url"`
	WebsiteOrigin string `gorm:"type:varchar(255); not null" json:"website_origin"`
	TotalReviews int `gorm:"not null" json:"total_reviews"`
}