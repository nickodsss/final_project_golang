package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null" validate:"required"`
	Email     string `gorm:"unique;not null" validate:"required,email"`
	Password  string `gorm:"not null" validate:"required,min=6"`
	Age       int    `gorm:"not null" validate:"required,min=9"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" validate:"required" json:"title"`
	Caption   string    `json:"caption,omitempty"`
	PhotoURL  string    `gorm:"not null" validate:"required" json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `gorm:"not null" validate:"required" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SocialMedia struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" validate:"required" json:"name"`
	SocialMediaURL string    `gorm:"not null" validate:"required" json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
