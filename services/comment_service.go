package services

import (
	"backend-platform/models"
	"backend-platform/repositories"
)

type CommentService interface {
	CreateComment(userID uint, discussionID uint, content string) (*models.Comment, error)
	// Add other methods as needed
}

type commentService struct {
	repo           repositories.CommentRepository
	discussionRepo repositories.DiscussionRepository
}

func NewCommentService(repo repositories.CommentRepository, discussionRepo repositories.DiscussionRepository) CommentService {
	return &commentService{repo, discussionRepo}
}

func (s *commentService) CreateComment(userID uint, discussionID uint, content string) (*models.Comment, error) {
	// Implement your business logic here, similar to the previous example
	comment := &models.Comment{
		UserID:       userID,
		DiscussionID: discussionID,
		Content:      content,
	}
	if err := s.repo.Create(comment); err != nil {
		return nil, err
	}
	return comment, nil
}
