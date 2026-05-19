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


	isMember, err := repository.IsChatMember(message.ChatID, message.SenderID)
	if err != nil {
		return err
	}
	if !isMember {
		return errors.New("not a chat member")
	}

	return repository.CreateMessage(message)
}

func GetChatMessages(chatID uint) ([]models.Message, error) {
	return repository.GetChatMessages(chatID)
}
