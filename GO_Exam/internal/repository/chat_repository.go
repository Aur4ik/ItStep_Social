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

func FindDMChat(
	user1 uint,
	user2 uint,
) (*models.Chat, error) {

	var chat models.Chat

	err := config.DB.
		Table("chats").
		Joins(
			"JOIN chat_members cm1 ON cm1.chat_id = chats.id",
		).
		Joins(
			"JOIN chat_members cm2 ON cm2.chat_id = chats.id",
		).
		Where("chats.is_dm = ?", true).
		Where("cm1.user_id = ?", user1).
		Where("cm2.user_id = ?", user2).
		First(&chat).Error

	if err != nil {
		return nil, err
	}

	return &chat, nil
}