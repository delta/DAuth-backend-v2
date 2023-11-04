package repository

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
)

type ResourceOwnerRepository interface {
	Insert(ctx context.Context, resource entity.ResourceOwner) entity.ResourceOwner
	Delete(ctx context.Context, resource entity.ResourceOwner)
	Exists(ctx context.Context, resource entity.ResourceOwner) bool
	FindByEmailID(ctx context.Context, resource int64) entity.ResourceOwner
	FindByID(ctx context.Context, resource int64) entity.ResourceOwner
}
