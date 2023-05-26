package entity


import (
	"gorm.io/gorm"
)

type ResourceOwner struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;column:id"`
	name	  string `gorm:"type:varchar(36);column:name"`
}

func (ResourceOwner) TableName() string {
	return "resource_owner"
}
