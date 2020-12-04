package main

import (
	"fmt"
	"golang_repository_restapi/controllers"
	"golang_repository_restapi/drivers"
	"golang_repository_restapi/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := drivers.ConnectDB()

	// repositories
	userRepository := repositories.NewUserRepository(db)

	// controller
	userController := controllers.NewUserController(userRepository)
	pingController := controllers.NewPingController()

	router := mux.NewRouter()
	router.HandleFunc("/ping", pingController.Ping).Methods("GET")
	router.HandleFunc("/users", userController.All).Methods("GET")

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s:%s", "localhost", "5000"),
	}
	server.ListenAndServe()

}
