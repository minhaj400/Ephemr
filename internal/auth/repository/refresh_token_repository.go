package repository

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/model"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(rt *model.RefreshTokens) error
}

type refreshTokenRepo struct {
	db *gorm.DB
}

func NewRefreshTokenRepo(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepo{db: db}
}

func (r *refreshTokenRepo) Create(rt *model.RefreshTokens) error {
	return r.db.Create(rt).Error
}
