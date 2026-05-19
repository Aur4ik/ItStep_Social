package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`

	AuthorID  uint      `json:"author_id"`
	Author    User      `json:"author" gorm:"foreignKey:AuthorID"`

	CommunityID *uint      `json:"community_id"`
	Community Community `json:"community" gorm:"constraint:OnDelete:CASCADE;"`

	LikesCount int64 `json:"likes_count" gorm:"-"`
	CommentsCount int64 `json:"comments_count" gorm:"-"`

	CreatedAt time.Time `json:"created_at"`
}