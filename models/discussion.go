package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Discussion struct {
	gorm.Model
	UserID    uint           `gorm:"not null"`
	Text      string         `gorm:"type:text"`
	Image     string         `gorm:"type:text"`
	HashTags  pq.StringArray `json:"hash_tags" gorm:"type:text[]"`
	CreatedOn string         `gorm:"not null"`
	Comments  []Comment
}
