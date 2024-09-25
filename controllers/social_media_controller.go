package controllers

import (
	"final_project_hacktiv8/config"
	"final_project_hacktiv8/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllSocialMedia(c *gin.Context) {
	var socialmedias []models.SocialMedia
	if err := config.DB.Find(&socialmedias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": socialmedias})
}

func GetSocialMedia(c *gin.Context) {
	var socialmedia models.SocialMedia
	id := c.Param("id")

	if err := config.DB.First(&socialmedia, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": socialmedia})
}

func CreateSocialMedia(c *gin.Context) {
	var input models.SocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.MustGet("userID").(jwt.MapClaims)

	input.UserID = uint(userID["id"].(float64))

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdateSocialMedia(c *gin.Context) {
	var socialmedia models.SocialMedia
	id := c.Param("id")

	if err := config.DB.First(&socialmedia, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social Media not found"})
		return
	}

	var input models.SocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&socialmedia).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": socialmedia})
}

func DeleteSocialMedia(c *gin.Context) {
	var socialmedia models.SocialMedia
	id := c.Param("id")

	if err := config.DB.First(&socialmedia, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social Media not found"})
		return
	}

	if err := config.DB.Delete(&socialmedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social Media deleted successfully"})
}
