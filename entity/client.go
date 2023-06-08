package entity

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ClientID        string          `gorm:"unique; column: client_id; primaryKey:false"`
	ClientSecret   string           `gorm:"type:varchar(191)"`
	Name           string           `gorm:"type:varchar(191)"`
	RedirectURI    string           `gorm:"type:varchar(191)"`
	Description    string           `gorm:"type:varchar(191)"`
	HomePageURI    string           `gorm:"type:varchar(191)"`
	UserID         int              `gorm:"not null"`
	Token          []Token          `gorm:"foreignKey:ClientID"`
	Code           []Code           `gorm:"foreignKey:ClientID"`
	AuthorisedApps []AuthorisedApps `gorm:"foreignKey:ClientID"`
}

func (Client) TableName() string {
	return "client"
}
