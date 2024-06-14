// services/discussion_service.go

package services

import (
	"backend-platform/models"
	"backend-platform/repositories"
)

type DiscussionService struct {
	repo *repositories.DiscussionRepository
}

func NewDiscussionService(repo *repositories.DiscussionRepository) *DiscussionService {
	return &DiscussionService{repo}
}

func (ds *DiscussionService) CreateDiscussion(discussion *models.Discussion) error {
	return ds.repo.CreateDiscussion(discussion)
}

func (ds *DiscussionService) UpdateDiscussion(discussion *models.Discussion) error {
	return ds.repo.UpdateDiscussion(discussion)
}
func (ds *DiscussionService) DeleteDiscussion(id uint) error {
	return ds.repo.DeleteDiscussion(id)
}
func (ds *DiscussionService) GetDiscussionsByTags(tags []string) ([]models.Discussion, error) {
	return ds.repo.GetDiscussionsByTags(tags)
}

func (ds *DiscussionService) GetDiscussionsBySearchText(searchText string) ([]models.Discussion, error) {
	return ds.repo.GetDiscussionsBySearchText(searchText)
}
