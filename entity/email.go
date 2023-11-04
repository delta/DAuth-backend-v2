package entity

import (
	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	ID             int64  `gorm:"primaryKey;auto_increment;not_null;column:id"`
	ActivationCode string `gorm:"unique"`
	Email          string `gorm:"unique"`
	IsActivated    bool   `gorm:"default:false"`
}

func (Email) TableName() string {
	return "email"
}
