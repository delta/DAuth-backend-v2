package entity

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model

	AccessToken string    `gorm:"type:varchar(191);not_null"`
	CreatedAt   time.Time `gorm:"not_null"`
	ExpireAt    time.Time `gorm:"not_null"`
	UserID      int64     `gorm:"not_null;column:user_id"`
	ClientID    int64     `gorm:"not_null;column:client_id"`
	Scope       string    `gorm:"type:varchar(191);not_null"`
}

func (Token) TableName() string {
	return "token"
}
