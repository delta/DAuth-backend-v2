package impl

import (
	"context"
	"errors"

	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/model"
	"github.com/delta/DAuth-backend-v2/repository"
	"github.com/delta/DAuth-backend-v2/service"
)

type emailServiceImpl struct {
	repository repository.EmailRepository
}

func NewEmailServiceImpl(rep repository.EmailRepository) service.EmailService {
	return &emailServiceImpl{rep}
}

func (impl *emailServiceImpl) FindByEmail(ctx context.Context, resource model.LoginRequest) (entity.Email, error) {
	var userEmail entity.Email

	if userEmail, _ = impl.repository.FindByEmail(ctx, resource); userEmail == (entity.Email{}) {
		return entity.Email{}, errors.New("Email Not Found")
	}
	return userEmail, nil
}
