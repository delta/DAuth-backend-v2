package entity

import (
	"gorm.io/gorm"
)

type Email struct{
	gorm.Model
	ActivationCode string     `gorm:"unique"`
	Email          string     `gorm:"unique"`
	IsActivated    bool       `gorm:"default:false"`
}

func (Email) TableName() string {
	return "email"
}
