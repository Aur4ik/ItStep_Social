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

func LeaveCommunity(userID uint, communityID uint) error {

	_, err := repository.GetMembership(userID, communityID)
	if err != nil {
		return errors.New("not a member")
	}

	return repository.DeleteMembership(userID, communityID)
}

func DeleteCommunity(
	communityID uint,
	userID uint,
	role string,
) error {

	community, err := repository.GetCommunityByID(communityID)
	if err != nil {
		return errors.New("community not found")
	}

	isOwner := community.OwnerID == userID

	isAdmin := role == "admin"

	if !isOwner && !isAdmin {
		return errors.New("forbidden")
	}

	return repository.DeleteCommunity(communityID)
}