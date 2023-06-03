package impl

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/repository"
	"github.com/delta/DAuth-backend-v2/service"
)

type authServiceImpl struct {
	repository repository.ResourceOwnerRepository
}

func NewAuthServiceImpl(rep repository.ResourceOwnerRepository) service.AuthService {
	return &authServiceImpl{rep}
}

func (impl *authServiceImpl) Create(ctx context.Context, resource entity.ResourceOwner) entity.ResourceOwner {
	return impl.repository.Insert(ctx, resource)
}

func (impl *authServiceImpl) Delete(ctx context.Context, resource entity.ResourceOwner) {
	if !impl.repository.Exists(ctx, resource) {
		return
	}

	impl.repository.Delete(ctx, resource)
}
