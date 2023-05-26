package impl

import (
	"context"
	"fmt"

	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/repository"
	"gorm.io/gorm"
)

func NewResourceOwnerRepositoryImpl (DB *gorm.DB) repository.ResourceOwnerRepository {
	return &resourceOwnerRepositoryImpl{DB: DB}
}

type resourceOwnerRepositoryImpl struct {
	*gorm.DB
}


func (repository *resourceOwnerRepositoryImpl) Insert(ctx context.Context, resource entity.ResourceOwner) entity.ResourceOwner {
	err := repository.DB.WithContext(ctx).Create(&resource).Error
	if err != nil {
		fmt.Println("Error")
	}
	return resource
}


func (repository *resourceOwnerRepositoryImpl) Delete(ctx context.Context, resource entity.ResourceOwner) {
	err := repository.DB.WithContext(ctx).Delete(&resource).Error
	if err != nil {
		fmt.Println("Error")
	}
}
