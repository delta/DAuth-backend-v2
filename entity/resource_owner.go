package entity

import (
	"gorm.io/gorm"
)

type ResourceOwner struct {
	gorm.Model
	Name           string `gorm:"type:varchar(191)"`
	Password       string `gorm:"type:varchar(191)"`
	Email          *Email `gorm:"foreignKey:EmailID"`
	EmailID        int
	PhoneNumber    string           `gorm:"type:varchar(191)"`
	Gender         Gender           `gorm:"default:MALE"`
	Batch          string           `gorm:"type:varchar(191)"`
	Token          []Token          `gorm:"foreignKey:UserID"`
	Code           []Code           `gorm:"foreignKey:UserID"`
	AuthorisedApps []AuthorisedApps `gorm:"foreignKey:UserID"`
	Client         []Client         `gorm:"foreignKey:UserID"`
}

type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

func (ResourceOwner) TableName() string {
	return "resource_owner"
}
