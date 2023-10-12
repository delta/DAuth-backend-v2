package database

import (
	"github.com/delta/DAuth-backend-v2/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	for _, entity := range []interface{}{
		entity.ResourceOwner{},
		entity.Email{},
		entity.Code{},
		entity.Client{},
		entity.Token{},
		entity.CodeChallenge{},
		entity.AuthorisedApps{},
	} {
		if err := db.AutoMigrate(&entity); err != nil {
			panic(err)
		}
	}
}
