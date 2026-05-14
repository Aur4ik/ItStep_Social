package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreatePost(post *models.Post) error{
	return config.DB.Create(post).Error
}
func GetPosts() ([]models.Post, error){
	var posts []models.Post

	err := config.DB.Preload("Author").Order("created_at desc").Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}
func GetPostByID(id uint) (*models.Post, error){
	var post models.Post

	err := config.DB.First(&post, id).Error

	if err != nil {
		return nil, err
	}
	return &post, nil
}

func DeletePost(id uint)error{
	return config.DB.Delete(&models.Post{}, id).Error
}

func GetPostWithAuthor(id uint) (*models.Post, error) {

	var post models.Post

	err := config.DB.
		Preload("Author").
		First(&post, id).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetCommunityPosts(communityID uint) ([]models.Post, error) {

	var posts []models.Post

	err := config.DB.
		Preload("Author").
		Where("community_id = ?", communityID).
		Order("created_at desc").
		Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}