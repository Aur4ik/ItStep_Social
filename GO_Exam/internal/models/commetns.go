package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`

	AuthorID  uint      `json:"author_id"`
	Author    User      `json:"author" gorm:"foreignKey:AuthorID"`

	PostID    uint      `json:"post_id"`

	CreatedAt time.Time `json:"created_at"`
}