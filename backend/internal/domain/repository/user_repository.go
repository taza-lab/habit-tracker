package repository

import (
	"backend/internal/domain/model/user"
)

type UserRepository interface {
	Find(id string) (*user.User, error)
	FindByUserName(username string) (*user.User, error)
	Register(user *user.User) (*user.User, error)
}
