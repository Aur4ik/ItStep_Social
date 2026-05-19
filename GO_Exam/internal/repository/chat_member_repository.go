package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func AddChatMember(member *models.ChatMember) error {
	return config.DB.Create(member).Error
}

func GetChatMembers(chatID uint) ([]models.User, error) {
	var users []models.User

	err := config.DB.
		Table("users").
		Select("users.*").
		Joins("JOIN chat_members ON chat_members.user_id = users.id").
		Where("chat_members.chat_id = ?", chatID).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}


func IsChatMember(chatID uint, userID uint) (bool, error) {
	var count int64
	err := config.DB.
		Model(&models.ChatMember{}).
		Where("chat_id = ? AND user_id = ?", chatID, userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
