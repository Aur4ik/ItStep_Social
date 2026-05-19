package models

import "time"

type Chat struct {
	ID uint `json:"id" gorm:"primaryKey"`

	IsDM bool `json:"is_dm"`

	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	
}