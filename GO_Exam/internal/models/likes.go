package models

type Like struct{
	ID uint `json:"id" gorm:"primaryKey"`

	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`
}