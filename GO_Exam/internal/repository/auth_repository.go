package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	err := config.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByID(id uint) (*models.User, error) {

	var user models.User

	err := config.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}