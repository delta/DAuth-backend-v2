package entity

import (
	"time"

	"gorm.io/gorm"
)

type AuthorisedApps struct {
	gorm.Model

	ClientID  string    `gorm:"type:varchar(191);not_null;column:client_id"`
	UserID    string    `gorm:"type:varchar(191);not_null;column:user_id"`
	CreatedAt time.Time `gorm:"not_null"`
}

func (AuthorisedApps) TableName() string {
	return "authorised_apps"
}
