package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/adhistria/auth-movie-app/internal/domain"
	mock_domain "github.com/adhistria/auth-movie-app/internal/domain/mock"
	"github.com/adhistria/auth-movie-app/internal/service"
	"github.com/golang/mock/gomock"
)

func TestUserService(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserRepository(ctrl)

	mockUser := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "password",
	}

	mu.EXPECT().Create(context.Background(), &mockUser).Return(nil)

	userService := service.UserService{
		UserRepo: mu,
	}

	err := userService.Register(context.Background(), &mockUser)
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}
}

func TestUserServiceWhenGeneratePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserRepository(ctrl)

	mockUser := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "password",
	}

	userService := service.UserService{
		UserRepo: mu,
	}

	oldHashAndSalt := service.HashAndSalt
	defer func() { service.HashAndSalt = oldHashAndSalt }()

	newHashAndSalt := func(password []byte, cost int) ([]byte, error) {
		return nil, errors.New("Error when generated password")
	}

	service.HashAndSalt = newHashAndSalt

	err := userService.Register(context.Background(), &mockUser)
	if err == nil {
		t.Errorf("Error was expected: %s", err)
	}
}
