package user

import (
	"net/http"

	"github.com/Prosp3r/userfarm/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// Get a user by ID
func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := h.DB.Preload("Tasks").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
