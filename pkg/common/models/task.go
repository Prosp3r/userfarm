package models

import (
	"time"

	"gorm.io/gorm"
)

// Task model
type Task struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Category    string `gorm:"not null"` // Category of the task (e.g., "Profile", "Content", "Purchase")
	Description string `gorm:"not null"`
	Points      int    `gorm:"not null"`
	Completed   bool   `gorm:"default:false"`
	CompletedAt *time.Time
}
