// repositories/discussion_repository.go

package repositories

import (
	"backend-platform/models"
	"strings"

	"gorm.io/gorm"
)

type DiscussionRepository struct {
	db *gorm.DB
}

func NewDiscussionRepository(db *gorm.DB) *DiscussionRepository {
	return &DiscussionRepository{db}
}

func (dr *DiscussionRepository) CreateDiscussion(discussion *models.Discussion) error {
	return dr.db.Create(discussion).Error
}

func (dr *DiscussionRepository) UpdateDiscussion(discussion *models.Discussion) error {
	return dr.db.Save(discussion).Error
}

func (dr *DiscussionRepository) DeleteDiscussion(id uint) error {
	// Soft delete if using gorm's soft delete feature
	return dr.db.Where("id = ?", id).Delete(&models.Discussion{}).Error

}

// GetDiscussionsByTags retrieves discussions containing the specified tags
func (dr *DiscussionRepository) GetDiscussionsByTags(tags []string) ([]models.Discussion, error) {
	var discussions []models.Discussion

	// Convert tags to lowercase for case-insensitive search
	for i := range tags {
		tags[i] = strings.ToLower(tags[i])
	}

	// Query discussions where any tag in hash_tags array matches the provided tags
	if err := dr.db.Where("array_to_string(hash_tags, '||') LIKE ?", "%"+strings.Join(tags, "%")+"%").Find(&discussions).Error; err != nil {
		return nil, err
	}

	return discussions, nil
}
func (dr *DiscussionRepository) GetDiscussionsBySearchText(searchText string) ([]models.Discussion, error) {
	var discussions []models.Discussion
	searchText = strings.ToLower(searchText)
	searchText = "%" + searchText + "%"

	if err := dr.db.Where("LOWER(text) LIKE ?", searchText).Find(&discussions).Error; err != nil {
		return nil, err
	}

	return discussions, nil
}
