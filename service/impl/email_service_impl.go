package impl

import (
	"context"

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

func (impl *emailServiceImpl) FindByEmail(ctx context.Context, resource model.LoginRequest) entity.Email {
	return impl.repository.FindByEmail(ctx, resource)
}
