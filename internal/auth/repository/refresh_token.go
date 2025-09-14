package repository

import (
	"errors"

	"github.com/Minhajxdd/Ephemr/internal/auth/model"
	"gorm.io/gorm"
)

// RefreshTokenRepository defines the methods for interacting with refresh tokens table in the database.
type RefreshTokenRepository interface {
	// Create a new refresh token in the database.
	Create(rt *model.RefreshTokens) error

	// Updates an existing refresh token record in the database.
	Update(rt *model.RefreshTokens) error

	// DeleteById removes a refresh token record from the database by its ID.
	DeleteById(id uint) error

	// Find Record With Arguments
	Find(rt *model.RefreshTokens) (*model.RefreshTokens, error)
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

func (r *refreshTokenRepo) Update(rt *model.RefreshTokens) error {
	return r.db.Save(rt).Error
}

func (r *refreshTokenRepo) DeleteById(id uint) error {
	result := r.db.Delete(&model.RefreshTokens{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return nil
	}

	return nil
}

func (r *refreshTokenRepo) Find(rt *model.RefreshTokens) (*model.RefreshTokens, error) {
	var t model.RefreshTokens
	if err := r.db.Where(rt).First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}
