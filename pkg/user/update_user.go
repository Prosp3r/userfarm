package user

import (
	"net/http"

	"github.com/Prosp3r/deepseek/userfarm/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// Update a user
func (h handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}
