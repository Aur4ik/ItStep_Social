package models

import "time"

type Chat struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
}