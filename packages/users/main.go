package main

import (
	"github.com/gorilla/mux"
	"github.com/lcnssantos/go-microservice/packages/users/controllers"
	"github.com/lcnssantos/go-microservice/packages/users/infra"
	"github.com/lcnssantos/go-microservice/packages/users/repository"
	"github.com/lcnssantos/go-microservice/packages/users/services"
	"github.com/lcnssantos/go-microservice/shared/middlewares"
	"log"
	"net/http"
)

func main() {
	db, err := infra.GetDatabaseConnection()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	hashService := services.NewHashService()
	userService := services.NewUserService(userRepository, hashService)
	userController := controllers.NewUserController(userService)

	router := mux.NewRouter()
	router.Use(middlewares.NewJsonMiddleware().Handler)

	router.HandleFunc("/user", userController.Create).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)
}
