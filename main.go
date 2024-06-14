package main

import (
	"backend-platform/config"
	"backend-platform/controllers"
	"backend-platform/repositories"
	"backend-platform/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	// // AutoMigrate models
	// db.AutoMigrate(&models.User{}, &models.Discussion{}, &models.Comment{}, &models.Reply{})

	// Setup repositories
	userRepo := repositories.NewUserRepository(db)
	discussionRepo := repositories.NewDiscussionRepository(db)
	commentRepo := repositories.NewCommentRepository(db)

	// Setup services
	userService := services.NewUserService(userRepo)
	discussionService := services.NewDiscussionService(discussionRepo)
	commentService := services.NewCommentService(commentRepo, *discussionRepo)

	// Setup controllers
	userController := controllers.NewUserController(userService)
	discussionController := controllers.NewDiscussionController(discussionService)
	commentController := controllers.NewCommentController(commentService)

	// Setup Gin router
	router := gin.Default()

	// Routes for users
	userRouter := router.Group("/users")
	{
		userRouter.POST("/login", userController.CreateUser)
		userRouter.PUT("/update/:id", userController.UpdateUser)
		userRouter.DELETE("/delete/:id", userController.DeleteUser)
		userRouter.GET("/", userController.GetUsers)
		userRouter.GET("/search", userController.SearchUsers)
		userRouter.POST("/follow", userController.FollowUser)
		userRouter.POST("/unfollow/:unfollowUserID", userController.UnfollowUser)
	}

	// Routes for discussions
	discussionRouter := router.Group("/discussions")
	{
		discussionRouter.POST("/create", discussionController.CreateDiscussion)
		discussionRouter.PUT("/update/:id", discussionController.UpdateDiscussion)
		discussionRouter.DELETE("/delete/:id", discussionController.DeleteDiscussion)
		discussionRouter.GET("/tags", discussionController.GetDiscussionsByTags) // Endpoint to get discussions by tags
		discussionRouter.GET("/search", discussionController.GetDiscussionsBySearchText)

	}

	commentRouter := router.Group("/comments")
	commentRouter.POST("/:discussionID", commentController.CreateComment)

	// Start server
	router.Run(":8080")
}
