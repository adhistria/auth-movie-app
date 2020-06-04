package user

import (
	"context"
)

type UserRepository interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) error
}