package user

import (
	"net/http"

	"github.com/Prosp3r/deepseek/userfarm/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// Create a new user
func (h handler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default values
	user.ProfileStatus = "active"
	user.AccessLevel = "basic"

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
