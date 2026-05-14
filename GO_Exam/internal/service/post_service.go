package service

import (
	"errors"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
)
func CreatePost(post *models.Post) error{

	if len(post.Content) == 0 {
		return errors.New("Post cannot be empty")
	}
	if len(post.Content) > 500 {
		return errors.New("Post cannot be over 500 symbols")
	}
	return repository.CreatePost(post)
}
func GetPosts() ([]models.Post, error){
	return repository.GetPosts()
}
func DeletePost(postID uint, userID uint) error{
	post, err := repository.GetPostByID(postID)
	if err != nil {
		return err
	}

	if post.AuthorID != userID {
		return errors.New("forbidden")
	}

	return repository.DeletePost(postID)
}
func GetCommunityPosts(communityID uint) ([]models.Post, error) {
	return repository.GetCommunityPosts(communityID)
}