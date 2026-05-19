package repository

import (
	"project/itStep/internal/config"
	"project/itStep/internal/models"
)

func CreatePost(post *models.Post) error {
	return config.DB.Create(post).Error
}


func GetPosts(limit, offset int) ([]models.Post, error) {
	var posts []models.Post

	err := config.DB.
		Preload("Author").
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}

	for i := range posts {
		var likesCount int64
		config.DB.Model(&models.Like{}).Where("post_id = ?", posts[i].ID).Count(&likesCount)
		posts[i].LikesCount = likesCount

		var commentsCount int64
		config.DB.Model(&models.Comment{}).Where("post_id = ?", posts[i].ID).Count(&commentsCount)
		posts[i].CommentsCount = commentsCount
	}

	return posts, nil
}

func GetPostByID(id uint) (*models.Post, error) {
	var post models.Post

	err := config.DB.Preload("Author").First(&post, id).Error
	if err != nil {
		return nil, err
	}

	var likesCount int64
	config.DB.Model(&models.Like{}).Where("post_id = ?", post.ID).Count(&likesCount)
	post.LikesCount = likesCount

	var commentsCount int64
	config.DB.Model(&models.Comment{}).Where("post_id = ?", post.ID).Count(&commentsCount)
	post.CommentsCount = commentsCount

	return &post, nil
}

func DeletePost(id uint) error {
	return config.DB.Delete(&models.Post{}, id).Error
}

func GetPostWithAuthor(id uint) (*models.Post, error) {
	var post models.Post
	err := config.DB.Preload("Author").First(&post, id).Error
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
