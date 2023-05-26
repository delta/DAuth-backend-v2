package service

import "context"

type AuthService interface {
	Login(ctx context.Context)
}