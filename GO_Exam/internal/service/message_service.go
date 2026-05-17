package service

import (
	"errors"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
)

func CreateMessage(message *models.Message) error {

	if len(message.Content) < 1 {
		return errors.New("empty message")
	}

	return repository.CreateMessage(message)
}

func GetChatMessages(chatID uint) ([]models.Message, error) {
	return repository.GetChatMessages(chatID)
}