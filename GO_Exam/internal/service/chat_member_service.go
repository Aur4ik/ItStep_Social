package service

import (
	"errors"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
)

func AddChatMember(
	chatID uint,
	userID uint,
) error {

	member := models.ChatMember{
		ChatID: chatID,
		UserID: userID,
	}

	return repository.AddChatMember(&member)
}

func GetChatMembers(chatID uint) ([]models.User, error) {

	users, err := repository.GetChatMembers(chatID)
	if err != nil {
		return nil, errors.New("failed to get members")
	}

	return users, nil
}