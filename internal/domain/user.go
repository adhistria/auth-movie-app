package domain

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
)

// User represent user entity
type User struct {
	ID       uint64 `db:"id"`
	Name     string `json:"name" db:"name" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
	RoleID   string `db:"role_id"`
}

// RegisterRequest struct represent json body when register
type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginRequest struct represent json body when login
type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserClaim ..
type UserClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// UserToken ..
type UserToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// UserService represent of user service
type UserService interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) (*UserToken, error)
}

// UserRepository represent interface of user repository
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, user *User) (*User, error)
}
