package entity

import (
	"gorm.io/gorm"
)

type ResourceOwner struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;auto_increment;not_null;column:id"`
	Name string `gorm:"type:varchar(36);column:name"`
}

func (ResourceOwner) TableName() string {
	return "resource_owner"
}
