package repository

import (
	"errors"

	"github.com/Minhajxdd/Ephemr/internal/auth/model"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(rt *model.RefreshTokens) error
	Update(rt *model.RefreshTokens) error
	DeleteById(id uint) error
	FindWithTokenDeviceIp(token, device, ipAddress string) (*model.RefreshTokens, error)
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

func (r *refreshTokenRepo) FindWithTokenDeviceIp(token, device, ipAddress string) (*model.RefreshTokens, error) {
	var rt model.RefreshTokens
	if err := r.db.Where(&model.RefreshTokens{
		TokenHash: token,
		Device:    device,
		IpAddress: ipAddress,
	}).First(&rt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rt, nil
}
