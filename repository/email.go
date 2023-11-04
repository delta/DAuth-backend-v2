package repository

import (
	"context"

	"github.com/delta/DAuth-backend-v2/entity"
	"github.com/delta/DAuth-backend-v2/model"
)

type EmailRepository interface {
	FindByEmail(ctx context.Context, resource model.LoginRequest) entity.Email
}
