package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowerID     uint   `gorm:"not null"` // User ID of the follower
	FollowingID    uint   `gorm:"not null"` // User ID of the user being followed
	FollowUserName string `gorm:"not null"` // Username of the user being followed
	User           User   `gorm:"foreignKey:FollowerID"`
	FollowUser     User   `gorm:"foreignKey:FollowingID"`
	// Add any other fields specific to the follow relationship here
}
