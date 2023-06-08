package entity

import (
	"gorm.io/gorm"
)

type Code struct {
	gorm.Model
	Code        string `gorm:"unique"`
	Scope       string `gorm:"null;type: varchar(191)"`
	State       string `gorm:"null; type: varchar(191)"`
	RedirectURI string `gorm:"type:varchar(191)"`
	Nonce       string `gorm:"null; type: varchar(191)"`
	UserID      int    `gorm:"not null"`
	ClientID    int    `gorm:"not null"`
}

func (Code) TableName() string {
	return "code"
}
