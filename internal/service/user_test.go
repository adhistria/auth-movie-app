package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/adhistria/auth-movie-app/internal/domain"
	mock_domain "github.com/adhistria/auth-movie-app/internal/domain/mock"
	. "github.com/adhistria/auth-movie-app/internal/service"
	"github.com/golang/mock/gomock"
)

func TestGivenDuplicateEmailWhenRegisterThenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserRepository(ctrl)
	mockUser := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "password",
	}
	mu.EXPECT().FindByEmail(context.Background(), &mockUser).Return(&mockUser, nil)

	userService := NewUserService(mu)

	err := userService.Register(context.Background(), &mockUser)
	if err == nil {
		t.Errorf("Error was expected: %s", err)
	}
}
func TestUserService(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserRepository(ctrl)

	mockUser := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "password",
	}

	mu.EXPECT().FindByEmail(context.Background(), &mockUser).Return(nil, nil)
	mu.EXPECT().Create(context.Background(), &mockUser).Return(nil)

	userService := NewUserService(mu)

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
	mu.EXPECT().FindByEmail(context.Background(), &mockUser).Return(nil, nil)

	userService := NewUserService(mu)

	oldHashAndSalt := HashAndSalt
	defer func() { HashAndSalt = oldHashAndSalt }()

	newHashAndSalt := func(password []byte, cost int) ([]byte, error) {
		return nil, errors.New("Error when generated password")
	}

	HashAndSalt = newHashAndSalt

	err := userService.Register(context.Background(), &mockUser)
	if err == nil {
		t.Errorf("Error was expected: %s", err)
	}
}

func TestGivenUnregisterUserWhenLoginThenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "password",
	}

	mu := mock_domain.NewMockUserRepository(ctrl)
	mu.EXPECT().FindByEmail(context.Background(), &mockUser).Return(nil, errors.New("User not found"))

	userService := NewUserService(mu)

	_, err := userService.Login(context.Background(), &mockUser)
	if err == nil {
		t.Errorf("Error was expected: %s", err)
	}
}

func TestGivenUncorrectPasswordWhenLoginThenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "password",
	}

	mockUser2 := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "12345",
	}

	mu := mock_domain.NewMockUserRepository(ctrl)
	mu.EXPECT().FindByEmail(context.Background(), &mockUser).Return(&mockUser2, nil)

	userService := NewUserService(mu)

	oldCompareHashPassword := CompareHashPassword
	defer func() { CompareHashPassword = oldCompareHashPassword }()

	newCompareHashPassword := func(hashPassword []byte, password []byte) error {
		return errors.New("Error compare hash password")
	}

	CompareHashPassword = newCompareHashPassword

	_, err := userService.Login(context.Background(), &mockUser)
	if err == nil {
		t.Errorf("Error was expected: %s", err)
	}
}

func TestGivenCorrectPasswordWhenLoginThenReturnSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "password",
	}

	mockUser2 := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "12345",
	}

	mu := mock_domain.NewMockUserRepository(ctrl)
	mu.EXPECT().FindByEmail(context.Background(), &mockUser).Return(&mockUser2, nil)

	userService := NewUserService(mu)

	oldCompareHashPassword := CompareHashPassword
	defer func() { CompareHashPassword = oldCompareHashPassword }()

	newCompareHashPassword := func(hashPassword []byte, password []byte) error {
		return nil
	}

	CompareHashPassword = newCompareHashPassword

	_, err := userService.Login(context.Background(), &mockUser)
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}
}

func TestGivenCorrectPasswordWhenLoginThenJWTClaim(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "password",
	}

	mockUser2 := domain.User{
		Name:     "john",
		Email:    "john@doe.com",
		Password: "12345",
	}

	mu := mock_domain.NewMockUserRepository(ctrl)
	mu.EXPECT().FindByEmail(context.Background(), &mockUser).Return(&mockUser2, nil)

	userService := NewUserService(mu)

	oldCompareHashPassword := CompareHashPassword
	defer func() { CompareHashPassword = oldCompareHashPassword }()

	newCompareHashPassword := func(hashPassword []byte, password []byte) error {
		return nil
	}

	CompareHashPassword = newCompareHashPassword
	_, err := userService.Login(context.Background(), &mockUser)
	if err != nil {
		t.Errorf("Error was expected: %s", err)
	}

}
