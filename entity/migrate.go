package entity

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	for _, entity := range[] interface {} {
		ResourceOwner{},
	} {
		if err := db.AutoMigrate(&entity); err != nil {
			panic(err)
		}
	}
}