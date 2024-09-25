package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"final_project_hacktiv8/models"
)

var DB *gorm.DB

func ConnectionDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=final_project_golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
}
