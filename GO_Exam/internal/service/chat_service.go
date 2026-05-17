package service

import (
	"errors"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
)

func CreateChat(chat *models.Chat) error {

	if len(chat.Name) < 2 {
		return errors.New("chat name too short")
	}

	return repository.CreateChat(chat)
}

func GetChats() ([]models.Chat, error) {
	return repository.GetChats()
}
func UpdateChat(
	chatID uint,
	name string,
) error {

	chat, err := repository.GetChatByID(chatID)
	if err != nil {
		return errors.New("chat not found")
	}

	if len(name) < 2 {
		return errors.New("chat name too short")
	}

	chat.Name = name

	return repository.UpdateChat(chat)
}

func DeleteChat(chatID uint) error {

	_, err := repository.GetChatByID(chatID)
	if err != nil {
		return errors.New("chat not found")
	}

	return repository.DeleteChat(chatID)
}