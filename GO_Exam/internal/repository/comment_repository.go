package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreateComment(comment *models.Comment) error{
	return config.DB.Create(comment).Error
}
func GetCommentByPostID(postID uint) ([]models.Comment, error) {

	var comments []models.Comment

	err := config.DB.
		Preload("Author").
		Where("post_id = ?", postID).
		Order("created_at desc").
		Find(&comments).Error

	if err != nil {
		return nil, err
	}
	return comments, nil
}