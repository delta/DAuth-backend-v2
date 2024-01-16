package entity

import (
	"time"

	"gorm.io/gorm"
)

type Code struct {
	gorm.Model

	ID            int64          `gorm:"primaryKey;auto_increment;not_null;column:id"`
	CodeChallenge *CodeChallenge `gorm:"foreignKey:CodeChallenge"`
	Code          string         `gorm:"type:varchar(191);not_null"`
	Scope         string         `gorm:"type:varchar(191)"`
	State         string         `gorm:"type:varchar(191)"`
	RedirectURI   string         `gorm:"type:varchar(191);not_null;column:redirect_uri"`
	Nonce         string         `gorm:"type:varchar(191)"`
	UserID        int64          `gorm:"not_null;column:user_id"`
	ClientID      int64          `gorm:"not_null;column:client_id"`
	CreatedAt     time.Time      `gorm:"not_null"`
	ExpireAt      time.Time      `gorm:"not_null"`
}

func (Code) TableName() string {
	return "code"
}
