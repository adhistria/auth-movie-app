package http

import (
	"encoding/json"
	"net/http"

	"github.com/adhistria/auth-movie-app/internal/domain"
	"github.com/adhistria/auth-movie-app/internal/validation"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// UserHandler represent http handler for user
type UserHandler struct {
	UserSerivce domain.UserService
	Validator   *validation.Validator
}

// Register add new user
func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Warnf("Error decode user body when register : %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	errors := u.Validator.Validate(user)
	if errors != nil {
		log.Warnf("Error validate register : %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errors)
		return
	}
	err = u.UserSerivce.Register(r.Context(), &user)
	if err != nil {
		log.Warnf("Error register user : %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := SuccessResponse{
		Message: "Success Register User",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

// Login authenticate user
func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Warnf("Error when decode json : %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := u.UserSerivce.Login(r.Context(), &user)
	if err != nil {
		log.Warnf("Email %s when login %s", user.Email, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := SuccessResponse{
		Message: "Login success",
		Data:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

// NewUserHandler return user handler
func NewUserHandler(router *httprouter.Router, userService domain.UserService, validator *validation.Validator) {
	userHandler := UserHandler{
		UserSerivce: userService,
		Validator:   validator,
	}
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
}
