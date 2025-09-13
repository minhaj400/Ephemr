package model

import (
	"time"

	"gorm.io/gorm"
)

type RefreshTokens struct {
	gorm.Model

	UserID    uint      `json:"user_id" gorm:"not null;constraint;OnDelete:CASCADE"`
	TokenHash string    `json:"token_hash" gorm:"type:text;index;not null"`
	Device    string    `json:"device" gorm:"type:text;not null"`
	IpAddress string    `json:"ip_address" gorm:"type:text;not null"`
	Revoked   bool      `json:"revoked" gorm:"default:false"`
	ExpiresAt time.Time `json:"expires_at" gorm:"index;not null"`
}
