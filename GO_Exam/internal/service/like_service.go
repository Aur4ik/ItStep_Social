package service

import (
	"errors"

	"gorm.io/gorm"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
)

func ToggleLike(userID uint, postID uint) (string, error) {

	like, err := repository.GetLike(userID, postID)


	if err == nil {

		err = repository.DeleteLike(like.ID)
		if err != nil {
			return "", err
		}

		return "like removed", nil
	}


	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	newLike := models.Like{
		UserID: userID,
		PostID: postID,
	}

	err = repository.CreateLike(&newLike)
	if err != nil {
		return "", err
	}

	return "like added", nil
}