package task

import (
	"net/http"
	"time"

	"github.com/Prosp3r/userfarm/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// Mark a task as completed
func (h handler) CompleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := h.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Update task as completed
	task.Completed = true
	now := time.Now()
	task.CompletedAt = &now

	// Award points to the user
	var user models.User
	h.DB.First(&user, task.UserID)
	user.Points += task.Points

	// Save changes
	h.DB.Save(&task)
	h.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Task completed", "points_awarded": task.Points})
}
