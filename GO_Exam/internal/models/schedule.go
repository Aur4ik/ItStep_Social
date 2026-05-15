package models

import "time"

type Schedule struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Date string `json:"date"`
	Time string `json:"time"`

	TeacherID uint `json:"teacher_id"`
	Teacher   User `json:"teacher" gorm:"foreignKey:TeacherID"`

	CommunityID uint      `json:"community_id"`
	Community   Community `json:"community"`

	CreatedAt time.Time `json:"created_at"`
}