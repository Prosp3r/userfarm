package models

import (
	"gorm.io/gorm"
)

// Referral model to track referrals and points
type Referral struct {
	gorm.Model
	ReferrerID   uint `gorm:"not null"`        // User who referred someone
	RefereeID    uint `gorm:"not null;unique"` // User who was referred
	PointsEarned int  `gorm:"not null"`        // Points earned by the referrer
}
