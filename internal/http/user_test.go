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

	NewUserHandler(router, mu)

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
