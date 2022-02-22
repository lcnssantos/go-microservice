package services

import "golang.org/x/crypto/bcrypt"

type HashService struct {
}

func NewHashService() *HashService {
	return &HashService{}
}

func (this HashService) Hash(input string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(password), err
}

func (this HashService) Compare(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}
