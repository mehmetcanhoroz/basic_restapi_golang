package controllers

import (
	"encoding/json"
	"fmt"
	"golang_repository_restapi/models"
	"net/http"
)

// UserControllerHandler will hold everything that controller needs
type UserControllerHandler struct {
	userRepository models.UserRepository
}

// NewUserController returns an instance of the controller handler
func NewUserController(userRepository models.UserRepository) *UserControllerHandler {
	return &UserControllerHandler{
		userRepository: userRepository,
	}
}

// All returns all records of users from db
func (handler *UserControllerHandler) All(w http.ResponseWriter, r *http.Request) {
	users, err := handler.userRepository.All()
	if err != nil {
		fmt.Println("Error", users)
		panic(error(err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
