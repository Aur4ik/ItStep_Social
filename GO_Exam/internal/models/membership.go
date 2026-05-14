package models

type Membership struct {
	ID uint `json:"id" gorm:"primaryKey"`

	UserID uint `json:"user_id"`
	User   User `gorm:"constraint:OnDelete:CASCADE;"`

	CommunityID uint      `json:"community_id"`
	Community   Community `gorm:"constraint:OnDelete:CASCADE;"`
}