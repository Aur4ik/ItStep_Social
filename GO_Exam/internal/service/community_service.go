package service

import (
	"errors"

	"gorm.io/gorm"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"

)

func CreateCommunity(community *models.Community) error {
	if len(community.Name) < 3 {
		return errors.New("community name too short")
	}
	return repository.CreateCommunity(community)
}


func GetCommunities() ([]models.Community, error) {
	return repository.GetCommunities()
}


func JoinCommunity(userID uint, communityID uint)(error){
	_, err := repository.GetMembership(userID, communityID)
	if err == nil {
		return errors.New("Alredy joined")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	membership := models.Membership{
		UserID: userID,
		CommunityID: communityID,
	}
	return repository.CreateMembership(&membership)
}

func GetCommunityByID(id uint) (*models.Community, error) {
	return repository.GetCommunityByID(id)
}

func GetCommunityMembers(communityID uint) ([]models.User, error) {
	return repository.GetCommunityMembers(communityID)
}