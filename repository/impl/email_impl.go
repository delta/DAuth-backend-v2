package impl

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/model"
	"github.com/delta/DAuth-backend-v2/repository"
	"gorm.io/gorm"
)

type emailRepositoryImpl struct {
	*gorm.DB
}

func NewEmailRepositoryImpl(DB *gorm.DB) repository.EmailRepository {
	return &emailRepositoryImpl{DB: DB}
}

func (repository *emailRepositoryImpl) FindByEmail(ctx context.Context, resource model.LoginRequest) (entity.Email, error) {
	var userEmail entity.Email
	err := repository.DB.WithContext(ctx).Where("email = ?", resource.Email).First(&userEmail).Error
	if err != nil {
		return entity.Email{}, err
	}
	return userEmail, nil
}
