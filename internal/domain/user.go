package domain

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
)

// User represent user entity
type User struct {
	ID       uint64 `db:"id"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
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
