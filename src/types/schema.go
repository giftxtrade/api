package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID uuid.UUID `gorm:"type:uuid; primary key" json:"id"`
	CreatedAt time.Time `gorm:"index; not null; default: now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"index; not null; default: now()" json:"updatedAt"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return nil
}

func (base *Base) BeforeUpdate(tx *gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil
}

type UserActionBase struct {
	CreatedById uuid.UUID `gorm:"type:uuid; index; not null" json:"-"`
	CreatedBy User `gorm:"foreignKey:CreatedById" json:"created_by"`
	ModifiedById uuid.UUID `gorm:"type:uuid; index; not null" json:"-"`
	ModifiedBy User `gorm:"foreignKey:ModifiedById" json:"modified_by"`
}

type User struct {
	Base
	Email string `gorm:"varchar(255); not null; index; unique" json:"email"`
	Name string `gorm:"varchar(255); not null" json:"name"`
	ImageUrl string `gorm:"varchar(255);" json:"imageUrl"`
	IsAdmin bool `gorm:"default: false" json:"-"`
	IsActive bool `gorm:"default: false" json:"isActive"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Base.BeforeCreate(tx)
	user.IsActive = true
	return nil
}

type Category struct {
	Base
	Name string `gorm:"type:varchar(30); not null; index; unique" json:"name"`
	Description string `gorm:"type:text; default: ''" json:"description"`
	Url string `gorm:"type:text" json:"url"`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	Base
	Title string `gorm:"type:text; not null; index" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	ProductKey string `gorm:"type:varchar(255); not null; index; unique" json:"productKey"`
	ImageUrl string `gorm:"type:text" json:"imageUrl"`
	Rating float32 `gorm:"type:float; not null; index" json:"rating"`
	Price float32 `gorm:"type:float(2); not null; index" json:"price"`
	OriginalUrl string `gorm:"type:text; not null" json:"originalUrl"`
	WebsiteOrigin string `gorm:"type:varchar(255); not null" json:"websiteOrigin"`
	TotalReviews uint `gorm:"not null" json:"totalReviews"`
	CategoryId uuid.UUID `gorm:"type:uuid; index" json:"-"`
	Category Category `gorm:"foreignKey:CategoryId" json:"category"`
}

type Event struct {
	Base
	UserActionBase
	Name string `gorm:"type:varchar(255); not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Budget float32 `gorm:"type:float(2); not null; index" json:"budget"`
	InviteMessage string `gorm:"type:text" json:"inviteMessage"`
	DrawAt time.Time `gorm:"index; not null" json:"drawAt"`
	CloseAt time.Time `gorm:"index; not null" json:"closeAt"`
	Slug string `gorm:"type:varchar(255); not null" json:"slug"`
	Participants []Participant `json:"participants,omitempty"`
}

type Participant struct {
	Base
	UserActionBase
	Email string `gorm:"type:varchar(255); not null; index; unique" json:"email"`
	Nickname string `gorm:"type:varchar(255)" json:"nickname"`
	Address string `gorm:"type:varchar(255)" json:"address"`
	Organizer bool `gorm:"type:boolean; default:false; not null" json:"organizer"`
	Participates bool `gorm:"type:boolean; default:true; not null" json:"participates"`
	Accepted bool `gorm:"type:boolean; default:false; not null" json:"accepted"`
	EventId uuid.UUID `gorm:"type:uuid; index; not null" json:"-"`
	Event Event `gorm:"foreignKey:EventId" json:"event"`
	UserId uuid.NullUUID `gorm:"type:uuid; index; default:null" json:"-"`
	User User `gorm:"foreignKey:UserId" json:"user"`
}