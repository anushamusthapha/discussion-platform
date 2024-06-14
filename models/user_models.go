package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string       `gorm:"not null"`
	MobileNo     string       `json:"mobile_no" gorm:"unique;not null"`
	Email        string       `gorm:"unique;not null"`
	PasswordHash string       `gorm:"not null"`
	AccessToken  string       `gorm:"-"`
	Discussions  []Discussion `gorm:"many2many:user_discussions;"`
	Followers    []*Follow    `gorm:"foreignKey:FollowingID"` // Slice of Follow structs where this user is followed
	Following    []*Follow    `gorm:"foreignKey:FollowerID"`
}
