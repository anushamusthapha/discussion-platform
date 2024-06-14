package controllers

import (
	"backend-platform/models"
	"backend-platform/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.service.CreateUser(&user); err != nil {
		switch err.Error() {
		case "user already exists with this email":
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
		case "user already exists with this mobile number":
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this mobile number"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		}
		return
	}
	c.Set("user_id", user.ID)
	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// logger.Log.Errorf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser.ID = uint(userID) // Set the user ID from the path parameter
	if err := uc.service.UpdateUser(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = uc.service.DeleteUser(uint(userID))
	if err != nil {
		if err == errors.New("user not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) SearchUsers(c *gin.Context) {
	query := c.Query("user_name")
	users, err := uc.service.SearchUsers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) FollowUser(c *gin.Context) {
	// Retrieve userID from header
	userIDStr := c.GetHeader("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind JSON request payload to struct
	var followRequest struct {
		FollowUserName string `json:"follow_user_name"`
	}
	if err := c.ShouldBindJSON(&followRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call UserService method to follow the user
	err = uc.service.FollowUser(uint(userID), followRequest.FollowUserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User followed successfully"})
}

func (uc *UserController) UnfollowUser(c *gin.Context) {
	userID := getUserIdFromContext(c) // Assuming you have a function to get current user ID from context
	unfollowUserID, err := strconv.ParseUint(c.Param("unfollowUserID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := uc.service.UnfollowUser(uint(userID), uint(unfollowUserID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully unfollowed user"})
}

// func (uc *UserController) CreateDiscussion(c *gin.Context) {
// 	userID := getUserIdFromContext(c)
// 	var discussion models.Discussion
// 	if err := c.ShouldBindJSON(&discussion); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	discussion.UserID = userID
// 	if err := uc.service.CreateDiscussion(&discussion); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, discussion)
// }

// func (uc *UserController) UpdateDiscussion(c *gin.Context) {
// 	userID := getUserIdFromContext(c)
// 	discussionID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
// 	var updatedDiscussion models.Discussion
// 	if err := c.ShouldBindJSON(&updatedDiscussion); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	updatedDiscussion.ID = uint(discussionID)
// 	updatedDiscussion.UserID = userID
// 	if err := uc.service.UpdateDiscussion(&updatedDiscussion); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, updatedDiscussion)
// }

// func (uc *UserController) DeleteDiscussion(c *gin.Context) {
// 	userID := getUserIdFromContext(c)
// 	discussionID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err := uc.service.DeleteDiscussion(userID, uint(discussionID)); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Discussion deleted successfully"})
// }

// func (uc *UserController) GetDiscussionsByTags(c *gin.Context) {
// 	tags := c.Query("tags")
// 	discussions, err := uc.service.GetDiscussionsByTags(tags)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, discussions)
// }

// func (uc *UserController) GetDiscussionsByText(c *gin.Context) {
// 	text := c.Query("text")
// 	discussions, err := uc.service.GetDiscussionsByText(text)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, discussions)
// }

// Helper function to get userID from context (for illustration purposes)
func getUserIdFromContext(c *gin.Context) uint {
	// Implement your logic to get user ID from token or session
	// This is a mock implementation
	userID, _ := c.Get("user_id")
	return userID.(uint)
}
