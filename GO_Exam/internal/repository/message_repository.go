package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreateMessage(message *models.Message) error {
	return config.DB.Create(message).Error
}

func GetChatMessages(chatID uint) ([]models.Message, error) {

	var messages []models.Message

	err := config.DB.
		Preload("Sender").
		Where("chat_id = ?", chatID).
		Order("created_at asc").
		Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}