package service

import (
	"errors"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
)

func CreateComment(comment *models.Comment) error{
	if len(comment.Content) == 0 {
		return errors.New("Comment cannot be empty")
	}
	if len(comment.Content) > 100 {
		return errors.New("Comment cannot be over 100 symbols")
	}

	return repository.CreateComment(comment)
}
func GetCommentByPostID(postID uint) ([]models.Comment, error) {
	return repository.GetCommentByPostID(postID)
}