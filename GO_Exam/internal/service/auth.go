package service

import (
	"errors"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func Register(user *models.User) error {

	if len(user.Password) < 6 {
		return errors.New("password too short")
	}

	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(hashed)
	user.Role = "user"

	return repository.CreateUser(user)
}

func Login(email, password string) (*models.User, error) {

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func GetUserByID(id uint) (*models.User, error) {
	return repository.GetUserByID(id)
}