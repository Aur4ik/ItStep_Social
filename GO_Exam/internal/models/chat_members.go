package models

type ChatMember struct {
	ID uint `json:"id" gorm:"primaryKey"`

	ChatID uint `json:"chat_id"`
	Chat   Chat `json:"chat"`

	UserID uint `json:"user_id"`
	User   User `json:"user"`
}