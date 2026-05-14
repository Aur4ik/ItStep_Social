package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreateCommunity(community *models.Community) error {
	return config.DB.Create(community).Error
}
func GetCommunities()([]models.Community, error){
	var communities []models.Community

	err := config.DB.
		Preload("Owner").
		Find(&communities).Error

	if err != nil {
		return nil, err
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