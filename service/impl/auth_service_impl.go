package impl

import (
	"context"

	"github.com/delta/DAuth-backend-v2/repository"
	"github.com/delta/DAuth-backend-v2/service"
)

type authServiceImpl struct{
	repository repository.ResourceOwnerRepository
}

func NewAuthRepositoryImpl(rep repository.ResourceOwnerRepository) service.AuthService{
	return &authServiceImpl{rep}
}

func (impl *authServiceImpl) Login(ctx context.Context) {

}
