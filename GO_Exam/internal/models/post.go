package models

type Post struct {
	ID       uint `gorm:"primaryKey"`
	Content  string

	AuthorID uint
	Author   User

}