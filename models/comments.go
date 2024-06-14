package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID       uint   `gorm:"not null"`
	DiscussionID uint   `gorm:"not null"`
	Text         string `gorm:"type:text"`
	Likes        int
	Replies      []Reply
}
