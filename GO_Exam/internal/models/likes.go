package models

type Like struct {
	ID uint `json:"id" gorm:"primaryKey"`

	UserID uint `json:"user_id" gorm:"uniqueIndex:idx_user_post_like"`
	PostID uint `json:"post_id" gorm:"uniqueIndex:idx_user_post_like"`
}
