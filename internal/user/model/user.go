package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `json:"name" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null;uniqueIndex"`
	Password    string    `json:"-"`
	IsVerified  bool      `json:"is_verified" gorm:"default:false"`
	LastLoginAt time.Time `json:"last_login_at" gorm:"autoUpdateTime"`
}
