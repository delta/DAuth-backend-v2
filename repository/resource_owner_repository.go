package repository

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
)

type ResourceOwnerRepository interface {
	Exists(ctx context.Context, resource entity.ResourceOwner) (bool, error)
	FindByEmailID(ctx context.Context, resource int64) (entity.ResourceOwner, error)
	FindByID(ctx context.Context, resource int64) (entity.ResourceOwner, error)
}
