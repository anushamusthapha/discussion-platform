// controllers/discussion_controller.go

package controllers

import (
	"backend-platform/models"
	"backend-platform/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiscussionController struct {
	service *services.DiscussionService
}

func NewDiscussionController(service *services.DiscussionService) *DiscussionController {
	return &DiscussionController{service}
}

func (dc *DiscussionController) CreateDiscussion(c *gin.Context) {
	var discussion models.Discussion
	if err := c.ShouldBindJSON(&discussion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dc.service.CreateDiscussion(&discussion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, discussion)
}

func (dc *DiscussionController) UpdateDiscussion(c *gin.Context) {
	// Get discussion ID from URL parameter
	discussionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discussion ID"})
		return
	}

	// Bind JSON data into discussion struct
	var discussion models.Discussion
	if err := c.ShouldBindJSON(&discussion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the ID of the discussion to update
	discussion.ID = uint(discussionID)

	// Call service to update the discussion
	if err := dc.service.UpdateDiscussion(&discussion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, discussion)
}

func (dc *DiscussionController) DeleteDiscussion(c *gin.Context) {
	// Get discussion ID from URL parameter
	discussionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discussion ID"})
		return
	}

	// Call service to delete the discussion
	if err := dc.service.DeleteDiscussion(uint(discussionID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discussion deleted successfully"})
}

func (dc *DiscussionController) GetDiscussionsByTags(c *gin.Context) {
	// Get tags from query parameters
	tags := c.QueryArray("tags")

	discussions, err := dc.service.GetDiscussionsByTags(tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, discussions)
}
func (dc *DiscussionController) GetDiscussionsBySearchText(c *gin.Context) {
	searchText := c.Query("q")
	discussions, err := dc.service.GetDiscussionsBySearchText(searchText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, discussions)
}
