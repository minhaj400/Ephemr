package repository

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/model"
	"gorm.io/gorm"
)

type EmailTokenRepository interface {
	Create(user *model.EmailToken) error
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
