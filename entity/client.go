package entity

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model

	ID             int64            `gorm:"primaryKey;auto_increment;not_null;column:id"`
	ClientID       string           `gorm:"type:varchar(191);not_null;column:client_id"`
	ClientSecret   string           `gorm:"type:varchar(191);not_null"`
	Name           string           `gorm:"type:varchar(191);not_null"`
	RedirectURI    string           `gorm:"type:varchar(191);not_null;column:redirect_uri"`
	UserID         int64            `gorm:"column:user_id;not_null"`
	CreatedAt      time.Time        `gorm:"not_null"`
	UpdatedAt      time.Time        `gorm:"not_null"`
	Description    string           `gorm:"type:varchar(191);not_null"`
	HomePageURL    string           `gorm:"type:varchar(191);not_null;column:home_page_url"`
	Code           []Code           `gorm:"foreignKey:UserID"`
	AuthorisedApps []AuthorisedApps `gorm:"foreignKey:UserID"`
	Token          []Token          `gorm:"foreignKey:UserID"`
}

func (Client) TableName() string {
	return "client"
}
