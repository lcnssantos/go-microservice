package services

import (
	"errors"
	"github.com/lcnssantos/go-microservice/packages/users/repository"
	"github.com/lcnssantos/go-microservice/packages/users/structs"
)

type UserService struct {
	userRepository *repository.UserRepository
	hashService    *HashService
}

func NewUserService(userRepository *repository.UserRepository, hashService *HashService) *UserService {
	return &UserService{userRepository: userRepository, hashService: hashService}
}

func (this UserService) Create(data *structs.CreateUser) error {
	if _, err := this.userRepository.FindOneByEmail(data.Email); err == nil {
		return errors.New("Email already exist")
	}

	hash, err := this.hashService.Hash(data.Password)

	if err != nil {
		return nil
	}

	data.Password = hash

	if err := this.userRepository.Create(data); err != nil {
		return err
	}

	return nil
}
