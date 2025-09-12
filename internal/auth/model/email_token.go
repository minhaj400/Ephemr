package model

import (
	"time"

	"gorm.io/gorm"
)

type TokenKind string

const (
	TokenKindVerify TokenKind = "verify"
	TokenKindReset  TokenKind = "reset"
)

type EmailToken struct {
	gorm.Model

	UserID    uint      `json:"user_id" gorm:"not null;index;constraint:OnDelete:CASCADE"`
	TokenHash string    `json:"token_hash" gorm:"type:text;not null"`
	Kind      TokenKind `json:"kind" gorm:"type:text;not null"`
	Used      bool      `json:"used" gorm:"default:false"`
	ExpiresAt time.Time `json:"expires_at" gorm:"index;not null"`
}
