package impl

import (
	"context"
	"fmt"

	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/repository"
	"gorm.io/gorm"
)

func NewResourceOwnerRepositoryImpl(DB *gorm.DB) repository.ResourceOwnerRepository {
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
		fmt.Printf("Error deleting, %s", err)
	}
}

func (repository *resourceOwnerRepositoryImpl) Exists(ctx context.Context, resource entity.ResourceOwner) bool {
	count := int64(0)
	err := repository.DB.WithContext(ctx).Find(&resource).Count(&count).Error

	if err != nil {
		fmt.Println("Error counting resource")
		return false
	}

	return count == 1
}

func (repository *resourceOwnerRepositoryImpl) FindByEmailID(ctx context.Context, resource int64) entity.ResourceOwner {
	var userDetails entity.ResourceOwner
	err := repository.DB.WithContext(ctx).Where("email_id = ?", resource).First(&userDetails).Error

	if err != nil {
		fmt.Println("Error : User not found")
	}

	return userDetails
}

func (repository *resourceOwnerRepositoryImpl) FindByID(ctx context.Context, resource int64) entity.ResourceOwner {
	var userDetails entity.ResourceOwner
	err := repository.DB.WithContext(ctx).Where("id = ?", resource).First(&userDetails).Error

	if err != nil {
		fmt.Println("Error : User not found")
	}

	return userDetails
}
