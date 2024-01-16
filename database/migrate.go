package database

import (
	"github.com/delta/DAuth-backend-v2/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	for _, entity := range []interface{}{
		entity.ResourceOwner{},
		entity.AuthorisedApps{},
		entity.Client{},
		entity.CodeChallenge{},
		entity.Code{},
		entity.Email{},
		entity.Token{},
	} {
		if err := db.AutoMigrate(&entity); err != nil {
			panic(err)
		}
	}
}
