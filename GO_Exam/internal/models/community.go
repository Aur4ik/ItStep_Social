package models

import "time"

type Community struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Avatar      string    `json:"avatar"`

	OwnerID     uint      `json:"owner_id"`
	Owner       User      `json:"owner" gorm:"foreignKey:OwnerID"`

	MembersCount int64 `json:"members_count" gorm:"-"`
	
	CreatedAt   time.Time `json:"created_at"`
}