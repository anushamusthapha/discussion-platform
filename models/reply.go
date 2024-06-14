package models

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	CommentID uint   `gorm:"not null"`
	Text      string `gorm:"type:text"`
	Likes     int
}
