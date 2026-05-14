package models

type Membership struct {
	ID          uint `json:"id" gorm:"primaryKey"`

	UserID      uint `json:"user_id"`
	CommunityID uint `json:"community_id"`
}