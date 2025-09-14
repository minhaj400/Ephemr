package repository

import (
	"errors"

	"github.com/Minhajxdd/Ephemr/internal/auth/model"
	"gorm.io/gorm"
)

// EmailTokenRepository defines the methods for interacting with email tokens table in the database.
type EmailTokenRepository interface {
	// Create persists a new email token.
	Create(user *model.EmailToken) error

	// IsValid checks if a given user ID and token hash correspond to a valid email token of a specific kind.
	// It returns the found EmailToken if valid, otherwise an error.
	IsValid(userId uint, tokenHash string, kind model.TokenKind) (*model.EmailToken, error)

	// DeleteById removes an email token from storage given its ID.
	DeleteById(id uint) error
}

type emailTokenRepo struct {
	db *gorm.DB
}

func NewEmailTokenRepository(db *gorm.DB) EmailTokenRepository {
	return &emailTokenRepo{db: db}
}

func (r *emailTokenRepo) Create(u *model.EmailToken) error {
	return r.db.Create(u).Error
}

func (r *emailTokenRepo) IsValid(userId uint, tokenHash string, kind model.TokenKind) (*model.EmailToken, error) {
	var t model.EmailToken
	if err := r.db.Where(&model.EmailToken{
		UserID:    userId,
		TokenHash: tokenHash,
		Kind:      kind,
		Used:      false,
	}).First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (r *emailTokenRepo) DeleteById(id uint) error {
	result := r.db.Delete(&model.EmailToken{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
