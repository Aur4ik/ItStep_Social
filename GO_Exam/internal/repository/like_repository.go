package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func GetLike(userID uint, postID uint)(*models.Like, error){
	var like models.Like

	err := config.DB.Where("user_id = ? AND post_id = ?", userID, postID).
	First(&like).Error

	if err != nil {
		return nil, err
	}
	
	return &like, nil
}

func CreateLike(like *models.Like)error{
	return config.DB.Create(like).Error
}

func DeleteLike(id uint) error {
	return config.DB.Delete(&models.Like{}, id).Error
}

func CountLikes(postID uint) (int64, error){
	var count int64

	err := config.DB.
	Model(&models.Like{}).
	Where("post_id = ?", postID).
	Count(&count).Error

	return count, err
}