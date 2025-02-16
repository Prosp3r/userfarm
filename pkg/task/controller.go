package task

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/task")
	// Task routes
	routes.POST("/", h.CreateTask)
	routes.PUT("/:id/complete", h.CompleteTask)
}
