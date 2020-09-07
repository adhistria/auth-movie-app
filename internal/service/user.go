package service

import (
	"context"
	"time"

	"github.com/adhistria/auth-movie-app/internal/domain"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultRefreshToken = 1 * time.Hour
	defaultAccessToken  = 15 * time.Minute
)

// HashAndSalt password
var HashAndSalt = bcrypt.GenerateFromPassword

// CompareHashPassword comparing password
var CompareHashPassword = bcrypt.CompareHashAndPassword

// NewWithClaims token
var NewWithClaims = jwt.NewWithClaims

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

func (s *userService) Login(ctx context.Context, user *domain.User) (*domain.UserToken, error) {
	fUser, err := s.UserRepo.FindByEmail(ctx, user)
	if err != nil {
		log.Infof("Can't find user by email : %s ", err)
		return nil, err
	}
	err = CompareHashPassword([]byte(fUser.Password), []byte(user.Password))
	if err != nil {
		log.Infof("Password not match : %s ", err)
		return nil, err
	}

	userToken := domain.UserToken{
		AccessToken:  s.generateToken(ctx, "access_token", user),
		RefreshToken: s.generateToken(ctx, "refresh_token", user),
	}

	return &userToken, nil
}

func (s *userService) generateToken(ctx context.Context, tokenType string, user *domain.User) string {
	userClaim := domain.UserClaim{
		user.Name,
		user.Email,
		jwt.StandardClaims{
			Subject:   tokenType,
			Issuer:    user.Email,
			ExpiresAt: time.Now().Add(defaultAccessToken).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	accessToken, err := token.SignedString([]byte("supec secret key"))
	if err != nil {
		log.Warnf("Error signed token : %s ", err)
		return ""
	}
	return accessToken
}

// NewUserService return object userService
func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{UserRepo: userRepo}
}
