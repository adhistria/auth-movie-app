package service

import (
	"context"

	"github.com/adhistria/auth-movie-app/internal/domain"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt password
var HashAndSalt = bcrypt.GenerateFromPassword

// CompareHashPassword comparing password
var CompareHashPassword = bcrypt.CompareHashAndPassword

// userService represent user servcie
type userService struct {
	UserRepo domain.UserRepository
}

// Register add new user
func (s *userService) Register(ctx context.Context, user *domain.User) error {
	hashPassword, err := HashAndSalt([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Warn("Fail when generate password")
		return err
	}

	user.Password = string(hashPassword)
	return s.UserRepo.Create(ctx, user)
}

func (s *userService) Login(ctx context.Context, user *domain.User) error {
	fUser, err := s.UserRepo.FindByEmail(ctx, user)
	if err != nil {
		log.Infof("Can't find user by email : %s ", err)
		return err
	}
	err = CompareHashPassword([]byte(fUser.Password), []byte(user.Password))
	if err != nil {
		log.Infof("Password not match : %s ", err)
		return err
	}
	// return jwt token
	return nil
}

// NewUserService return object userService
func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{UserRepo: userRepo}
}
