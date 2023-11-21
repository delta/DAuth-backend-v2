package impl

import (
	"context"

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

func (repository *resourceOwnerRepositoryImpl) Exists(ctx context.Context, resource entity.ResourceOwner) (bool, error) {
	count := int64(0)
	err := repository.DB.WithContext(ctx).Find(&resource).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func (repository *resourceOwnerRepositoryImpl) FindByEmailID(ctx context.Context, resource int64) (entity.ResourceOwner, error) {
	var userDetails entity.ResourceOwner
	err := repository.DB.WithContext(ctx).Where("email_id = ?", resource).First(&userDetails).Error

	if err != nil {
		return entity.ResourceOwner{}, err
	}

	return userDetails, nil
}

func (repository *resourceOwnerRepositoryImpl) FindByID(ctx context.Context, resource int64) (entity.ResourceOwner, error) {
	var userDetails entity.ResourceOwner
	err := repository.DB.WithContext(ctx).Where("id = ?", resource).First(&userDetails).Error

	if err != nil {
		return entity.ResourceOwner{}, err
	}

	return userDetails, nil
}
