package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adhistria/auth-movie-app/internal/domain"
	mock_domain "github.com/adhistria/auth-movie-app/internal/domain/mock"
	. "github.com/adhistria/auth-movie-app/internal/http"
	"github.com/adhistria/auth-movie-app/internal/validation"
	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
)

func TestWhenRegisterGivenIncorectUserThenReturnError(t *testing.T) {

	user := domain.User{
		Name:     "adhi",
		Email:    "adhistria1@gmail.com",
		Password: "password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserService(ctrl)
	mu.EXPECT().Register(context.Background(), &user).Return(errors.New("Error create new user"))
	router := httprouter.New()
	validator := validation.NewValidator()

	NewUserHandler(router, mu, validator)

	registerBody, err := json.Marshal(map[string]interface{}{
		"email":    "adhistria1@gmail.com",
		"name":     "adhi",
		"password": "password",
	})
	if err != nil {
		t.Fatalf("Fail when marshal JSON : %s", err)
	}

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(registerBody))
	if err != nil {
		t.Fatalf("Error create new request : %s", err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code")
	}
	t.Logf("Status Code : %v", rr.Code)
	t.Logf("Status Code : %v", rr.Result())
}

func TestWhenRegisterGivenIncorrectBodyThenReturnError(t *testing.T) {

	router := httprouter.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserService(ctrl)
	validator := validation.NewValidator()
	NewUserHandler(router, mu, validator)

	registerBody, err := json.Marshal(map[string]interface{}{
		"email":    1,
		"name":     2,
		"password": 3,
	})
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(registerBody))
	if err != nil {
		t.Fatalf("Error create new request : %s", err)
	}

	rr := httptest.NewRecorder()
	t.Log(rr)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code")
	}
	t.Logf("Status Code : %v", rr.Code)
}

func TestGiveIncorrectBodyWhenRegisterThenReturnError(t *testing.T) {
	router := httprouter.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserService(ctrl)
	validator := validation.NewValidator()
	NewUserHandler(router, mu, validator)

	registerBody, err := json.Marshal(map[string]interface{}{
		"email":    "",
		"name":     "",
		"password": "",
	})
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(registerBody))
	if err != nil {
		t.Fatalf("Error create new request : %s", err)
	}

	rr := httptest.NewRecorder()
	t.Log(rr)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code")
	}
	t.Log("YYANG DISINI")
	t.Logf("Status Code : %v", rr.Code)
}

func TestWhenRegisterGivenCorrectUserThenReturnSuccess(t *testing.T) {
	user := domain.User{
		Name:     "adhi",
		Email:    "adhistria1@gmail.com",
		Password: "password",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserService(ctrl)
	mu.EXPECT().Register(context.Background(), &user).Return(nil)

	router := httprouter.New()
	validator := validation.NewValidator()
	NewUserHandler(router, mu, validator)

	registerBody, err := json.Marshal(map[string]interface{}{
		"email":    "adhistria1@gmail.com",
		"name":     "adhi",
		"password": "password",
	})
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(registerBody))
	if err != nil {
		t.Fatalf("Error create new request : %s", err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code")
	}
	t.Logf("Status Code : %v", rr.Code)
}

func TestWhenLoginGivenInvalidBodyThenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserService(ctrl)
	router := httprouter.New()
	validator := validation.NewValidator()
	NewUserHandler(router, mu, validator)

	invalidBody := map[string]interface{}{
		"Name":     1,
		"Email":    1,
		"Password": 1,
	}
	ibReq, err := json.Marshal(invalidBody)
	if err != nil {
		t.Fatalf("Error when marshal : %s", err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(ibReq))
	if err != nil {
		t.Fatalf("Error when create request : %s", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code")
	}
	t.Logf("Status code : %v ", rr.Code)
}

func TestLoginGivenInvalidUserWhenLoginThenReturnError(t *testing.T) {
	body := domain.User{
		Email:    "adhistria1@gmail.com",
		Password: "password",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserService(ctrl)
	mu.EXPECT().Login(context.Background(), &body).Return(nil, errors.New("Error when find user"))
	router := httprouter.New()
	validator := validation.NewValidator()
	NewUserHandler(router, mu, validator)

	userReq, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Fail marshal data : %s", err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userReq))
	if err != nil {
		t.Errorf("Fail create request : %s", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler return invalid status code")
	}
	t.Logf("Status code : %v ", rr.Code)

}

func TestGivenValidUserWhenLoginThenSuccess(t *testing.T) {
	body := domain.User{
		Email:    "adhistria1@gmail.com",
		Password: "password",
	}
	userToken := domain.UserToken{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserService(ctrl)
	mu.EXPECT().Login(context.Background(), &body).Return(&userToken, nil)
	router := httprouter.New()
	validator := validation.NewValidator()
	NewUserHandler(router, mu, validator)

	userReq, err := json.Marshal(body)
	if err != nil {
		t.Errorf("Fail marshal data : %s", err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userReq))
	if err != nil {
		t.Errorf("Fail create request : %s", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return invalid status code")
	}
	t.Logf("Status code : %v ", rr.Code)
}

func TestGivenInvalidBodyWhenLoginThenReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := mock_domain.NewMockUserService(ctrl)
	router := httprouter.New()
	validator := validation.NewValidator()
	NewUserHandler(router, mu, validator)

	invalidBody := map[string]interface{}{
		"Email":    "",
		"Password": "",
	}

	ibReq, err := json.Marshal(invalidBody)
	if err != nil {
		t.Fatalf("Error when marshal : %s", err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(ibReq))
	if err != nil {
		t.Fatalf("Error when create request : %s", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler return wrong status code")
	}
	t.Logf("Status code : %v ", rr.Code)
}
