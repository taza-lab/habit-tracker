package repository

import (
	"backend/internal/domain/model/user"
	"context"
)

type UserRepository interface {
	Find(ctx context.Context, id string) (*user.User, error)
	FindByUserName(ctx context.Context, username string) (*user.User, error)
	Register(ctx context.Context, user *user.User) (*user.User, error)
	UpdatePoints(ctx context.Context, userId string, points int) error
}
