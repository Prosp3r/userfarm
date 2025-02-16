package user

import (
	"net/http"

	"github.com/Prosp3r/deepseek/userfarm/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// Delete a user
func (h handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
