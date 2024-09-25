package main

import (
	"final_project_hacktiv8/config"
	"final_project_hacktiv8/controllers"
	"final_project_hacktiv8/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectionDatabase()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	protected := router.Group("/")
	protected.Use(middleware.Authentication()) // All routes here require authentication
	{
		// Photo Routes
		protected.GET("/photos", controllers.GetAllPhotos)
		protected.GET("/photos/:id", controllers.GetPhoto)
		protected.POST("/photos", controllers.CreatePhoto)
		protected.PUT("/photos/:id", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
		protected.DELETE("/photos/:id", middleware.PhotoAuthorization(), controllers.DeletePhoto)

		// Comment Routes
		protected.GET("/comments", controllers.GetAllComments)
		protected.GET("/comments/:id", controllers.GetComment)
		protected.POST("/comments", controllers.CreateComment)
		protected.PUT("/comments/:id", middleware.CommentAuthorization(), controllers.UpdateComment)
		protected.DELETE("/comments/:id", middleware.CommentAuthorization(), controllers.DeleteComment)

		// Social Media Routes
		protected.GET("/socialmedia", controllers.GetAllSocialMedia)
		protected.GET("/socialmedia/:id", controllers.GetSocialMedia)
		protected.POST("/socialmedia", controllers.CreateSocialMedia)
		protected.PUT("/socialmedia/:id", middleware.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		protected.DELETE("/socialmedia/:id", middleware.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	router.Run(":8080")
}
