package http

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/adhistria/auth-movie-app/internal/domain"
	log "github.com/sirupsen/logrus"
)

// UserHandler represent http handler for user
type UserHandler struct {
	UserSerivce domain.UserService
}

// Register add new user
func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Warnf("Error validate user body when register : %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		log.Warnf("User with email %s error when login %s", user.Email, err)
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

func NewUserHandler(router *httprouter.Router, userService domain.UserService) {
	userHandler := UserHandler{
		UserSerivce: userService,
	}
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
}
