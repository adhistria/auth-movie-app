package service

import (
	"context"

	"github.com/adhistria/auth-movie-app/internal/domain"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var HashAndSalt = bcrypt.GenerateFromPassword

// UserService represent user servcie
type UserService struct {
	UserRepo domain.UserRepository
}

// Register add new user
func (s *UserService) Register(ctx context.Context, user *domain.User) error {
	hashPassword, err := HashAndSalt([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Warn("Fail when generate password")
		return err
	}

	user.Password = string(hashPassword)
	return s.UserRepo.Create(ctx, user)
}
