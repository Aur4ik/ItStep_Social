package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreateChat(chat *models.Chat) error{
	return config.DB.Create(chat).Error
}
func GetChats()([]models.Chat, error){
	var chats []models.Chat

	err := config.DB.Find(&chats).Error

	if err != nil {
		return nil, err
	}
	return chats, nil
}
func GetChatByID(id uint) (*models.Chat, error) {

	var chat models.Chat

	err := config.DB.
		First(&chat, id).Error

	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func UpdateChat(chat *models.Chat) error {
	return config.DB.Save(chat).Error
}

func DeleteChat(id uint) error {

	return config.DB.
		Delete(&models.Chat{}, id).Error
}