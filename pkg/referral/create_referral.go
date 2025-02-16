package referral

import (
	"net/http"

	"github.com/Prosp3r/userfarm/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// Create a referral
func (h handler) CreateReferral(c *gin.Context) {
	var referral models.Referral
	if err := c.ShouldBindJSON(&referral); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate points for the referrer (10% of referee's points)
	var referee models.User
	h.DB.First(&referee, referral.RefereeID)
	pointsEarned := int(float64(referee.Points) * 0.10)
	referral.PointsEarned = pointsEarned

	// Award points to the referrer
	var referrer models.User
	h.DB.First(&referrer, referral.ReferrerID)
	referrer.Points += pointsEarned

	// Save referral and update referrer's points
	h.DB.Create(&referral)
	h.DB.Save(&referrer)

	c.JSON(http.StatusCreated, referral)
}
