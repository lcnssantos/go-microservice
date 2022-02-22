package controllers

import (
	"github.com/lcnssantos/go-microservice/packages/users/services"
	"github.com/lcnssantos/go-microservice/packages/users/structs"
	"github.com/lcnssantos/go-microservice/shared/utils"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{userService: service}
}

func (this UserController) Create(w http.ResponseWriter, r *http.Request) {
	createUser := &structs.CreateUser{}

	if err := utils.HandleValidateRequest(w, r, createUser); err != nil {
		return
	}

	if err := this.userService.Create(createUser); err != nil {
		utils.ThrowHttpError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
