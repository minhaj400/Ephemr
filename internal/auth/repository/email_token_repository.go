package repository

import (
	"errors"

	"github.com/Minhajxdd/Ephemr/internal/auth/model"
	"gorm.io/gorm"
)

type EmailTokenRepository interface {
	Create(user *model.EmailToken) error
	IsValid(userId uint, tokenHash string, kind model.TokenKind) (*model.EmailToken, error)
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
	err := r.db.Delete(&model.EmailToken{}, id)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return err.Error
}
