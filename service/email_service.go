package service

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/model"
)

type EmailService interface {
	FindByEmail(ctx context.Context, resource model.LoginRequest) (entity.Email, error)
}
