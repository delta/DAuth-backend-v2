package entity

import (
	"gorm.io/gorm"
)

type AuthorisedApps struct {
	gorm.Model
	ClientID int `gorm:"not null"`
	UserID   int `gorm:"not null"`
}

func (AuthorisedApps) TableName() string {
	return "authorised_apps"
}
