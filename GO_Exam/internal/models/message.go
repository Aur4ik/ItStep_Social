package models

import "time"

type Message struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Content string `json:"content"`

	ChatID uint `json:"chat_id"`
	Chat   Chat `json:"chat"`

	SenderID uint `json:"sender_id"`
	Sender   User `json:"sender"`

	CreatedAt time.Time `json:"created_at"`
}