package service

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
)

type ResourceService interface {
	FindByEmailID(ctx context.Context, resource int64) (entity.ResourceOwner, error)
	FindByID(ctx context.Context, resource int64) (entity.ResourceOwner, error)
}
