package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`

	AuthorID  uint      `json:"author_id"`
	Author    User      `json:"author" gorm:"foreignKey:AuthorID"`

	CommunityID *uint      `json:"community_id"`
	Community   Community  `json:"community"`

	CreatedAt time.Time `json:"created_at"`
}