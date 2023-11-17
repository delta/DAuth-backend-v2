package entity

import (
	"time"

	"gorm.io/gorm"
)

type CodeChallenge struct {
	gorm.Model

	ID                  int64     `gorm:"primaryKey;auto_increment;not_null;column:id"`
	CodeChallenge       string    `gorm:"type:varchar(191);not_null"`
	CodeChallengeMethod string    `gorm:"type:varchar(191);not_null"`
	CodeID              int64     `gorm:"not_null;column:code_id"`
	CreatedAt           time.Time `gorm:"not_null"`
	ExpireAt            time.Time `gorm:"not_null"`
}

func (CodeChallenge) TableName() string {
	return "code_challenge"
}
