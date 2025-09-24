package service

import (
	userModel "backend/internal/domain/model/user"
	"context"
)

type UserService interface {
	SignUp(ctx context.Context, userName string, password string) (*userModel.User, error)
	Login(ctx context.Context, userName string, password string) (*userModel.User, string, error)
}
