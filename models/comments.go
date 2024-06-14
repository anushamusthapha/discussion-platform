package models

import "gorm.io/gorm"

// Comment model
type Comment struct {
	gorm.Model
	Content      string `gorm:"not null"`
	UserID       uint   `gorm:"not null"`
	User         User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DiscussionID uint   `gorm:"not null"`
	Discussion   Discussion
}
