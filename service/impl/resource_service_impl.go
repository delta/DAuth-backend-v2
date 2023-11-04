package impl

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/repository"
	"github.com/delta/DAuth-backend-v2/service"
)

type resourceServiceImpl struct {
	repository repository.ResourceOwnerRepository
}

func NewResourceServiceImpl(rep repository.ResourceOwnerRepository) service.ResourceService {
	return &resourceServiceImpl{rep}
}

func (impl *resourceServiceImpl) Create(ctx context.Context, resource entity.ResourceOwner) entity.ResourceOwner {
	return impl.repository.Insert(ctx, resource)
}

func (impl *resourceServiceImpl) Delete(ctx context.Context, resource entity.ResourceOwner) {
	if !impl.repository.Exists(ctx, resource) {
		return
	}

	impl.repository.Delete(ctx, resource)
}

func (impl *resourceServiceImpl) FindByEmailID(ctx context.Context, resource int64) entity.ResourceOwner {
	return impl.repository.FindByEmailID(ctx, resource)
}

func (impl *resourceServiceImpl) FindByID(ctx context.Context, resource int64) entity.ResourceOwner {
	return impl.repository.FindByID(ctx, resource)
}
