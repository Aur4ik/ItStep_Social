package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreateCommunity(community *models.Community) error {
	return config.DB.Create(community).Error
}
func GetCommunities() ([]models.Community, error) {

	var communities []models.Community

	err := config.DB.
		Preload("Owner").
		Find(&communities).Error

	if err != nil {
		return nil, err
	}

	for i := range communities {

		var count int64

		config.DB.
			Model(&models.Membership{}).
			Where("community_id = ?", communities[i].ID).
			Count(&count)

		communities[i].MembersCount = count
	}

	return communities, nil
}
func CreateMembership(membership *models.Membership)error{
	return config.DB.Create(membership).Error
}
func GetMembership(userID uint, communityID uint)(*models.Membership, error){
	var membership models.Membership

	err := config.DB.
	Where("user_id = ? AND community_id = ?", userID, communityID).
	First(&membership).Error

	if err != nil {
		return nil, err
	}
	return &membership, nil
}
func GetCommunityByID(id uint) (*models.Community, error) {

	var community models.Community

	err := config.DB.
		Preload("Owner").
		First(&community, id).Error

	if err != nil {
		return nil, err
	}

	return &community, nil
}
func GetCommunityMembers(communityID uint) ([]models.User, error) {

	var users []models.User

	err := config.DB.
		Table("users").
		Select("users.*").
		Joins("JOIN memberships ON memberships.user_id = users.id").
		Where("memberships.community_id = ?", communityID).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteMembership(userID uint, communityID uint) error {

	return config.DB.
		Where("user_id = ? AND community_id = ?", userID, communityID).
		Delete(&models.Membership{}).Error
}

func DeleteCommunity(id uint) error {

	return config.DB.
		Delete(&models.Community{}, id).Error
}