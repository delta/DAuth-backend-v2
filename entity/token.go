package entity

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	AccessToken string `gorm:"type:varchar(191);unique"`
	UserID      int    `gorm:"not null"`
	ClientID    int    `gorm:"not null"`
	Scope       string `gorm:"type:varchar(191)"`
}

func (Token) TableName() string {
	return "token"
}
