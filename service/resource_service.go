package service

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
)

type ResourceService interface {
	Create(ctx context.Context, resource entity.ResourceOwner) entity.ResourceOwner
	Delete(ctx context.Context, resource entity.ResourceOwner)
	FindByEmailID(ctx context.Context, resource int64) entity.ResourceOwner
	FindByID(ctx context.Context, resource int64) entity.ResourceOwner
}
