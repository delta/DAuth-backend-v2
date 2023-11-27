package impl

import (
	"context"
	"errors"
	"reflect"

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

func (impl *resourceServiceImpl) FindByEmailID(ctx context.Context, resource int64) (entity.ResourceOwner, error) {
	var userDetails entity.ResourceOwner

	if userDetails, _ = impl.repository.FindByEmailID(ctx, resource); reflect.DeepEqual(userDetails, entity.ResourceOwner{}) {
		return entity.ResourceOwner{}, errors.New("User Not Found")
	}
	return userDetails, nil
}

func (impl *resourceServiceImpl) FindByID(ctx context.Context, resource int64) (entity.ResourceOwner, error) {
	var userDetails entity.ResourceOwner

	if userDetails, _ = impl.repository.FindByID(ctx, resource); reflect.DeepEqual(userDetails, entity.ResourceOwner{}) {
		return entity.ResourceOwner{}, errors.New("User Not Found")
	}
	return userDetails, nil
}
