package service

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
)

type AuthService interface {
	Create(ctx context.Context, resource entity.ResourceOwner) entity.ResourceOwner
	Delete(ctx context.Context, resource entity.ResourceOwner)
}