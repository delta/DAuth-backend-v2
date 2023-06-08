package entity

import (
	"gorm.io/gorm"
)

type CodeChallenge struct {
	gorm.Model
	CodeChallenge       string `gorm:"unique"`
	CodeChallengeMethod string `gorm:"type:varchar(191)"`
	Code                *Code  `gorm:"foreignKey:CodeID"`
	CodeID              int
}

func (CodeChallenge) TableName() string {
	return "code_challenge"
}
