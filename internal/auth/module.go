package auth

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/controller"
	"gorm.io/gorm"
)

type Module struct {
	Controller *controller.AuthController
}

func NewModule(db *gorm.DB) {
	
}