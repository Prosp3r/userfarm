package models

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Email          string `gorm:"unique;not null"`
	PasswordHash   string `gorm:"not null"`
	Is2FAEnabled   bool   `gorm:"default:false"`
	MetaMaskWallet string // Optional Crypto authentication
	ProfileStatus  string `gorm:"default:'active'"` // active or inactive
	AccessLevel    string `gorm:"default:'basic'"`  // basic or admin
	Points         int    `gorm:"default:0"`
	ReferralCode   string `gorm:"unique"`
	ReferredBy     *uint  // ID of the user who referred this user
	Tasks          []Task `gorm:"foreignKey:UserID"`
}
