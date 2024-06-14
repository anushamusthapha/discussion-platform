package controllers

import (
	"backend-platform/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) *CommentController {
	return &CommentController{service}
}

func (c *CommentController) CreateComment(ctx *gin.Context) {
	// Parse discussionID from path parameter
	discussionID, err := strconv.ParseUint(ctx.Param("discussionID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discussion ID"})
		return
	}

	// Parse request body
	var req CreateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Get userID from authentication middleware or session
	userID := uint(1) // Replace with actual authentication logic

	// Call service to create comment
	comment, err := c.service.CreateComment(userID, uint(discussionID), req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}
