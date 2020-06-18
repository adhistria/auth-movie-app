package domain

import "context"

// User represent user entity
type User struct {
	ID       uint64
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

// UserService represent of user service
type UserService interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) error
}

// UserRepository represent interface of user repository
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) error
}
