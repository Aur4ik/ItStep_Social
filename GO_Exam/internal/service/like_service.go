package service

import (
	"errors"

	"gorm.io/gorm"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
)


func ToggleLike(userID uint, postID uint) (string, int64, error) {

	like, err := repository.GetLike(userID, postID)

	if err == nil {
		if err = repository.DeleteLike(like.ID); err != nil {
			return "", 0, err
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		newLike := models.Like{UserID: userID, PostID: postID}
		if err = repository.CreateLike(&newLike); err != nil {
			return "", 0, err
		}
	} else {
		return "", 0, err
	}

	count, err := repository.CountLikes(postID)
	if err != nil {
		return "", 0, err
	}

	message := "like added"
	if like != nil {
		message = "like removed"
	}

	return message, count, nil
}
