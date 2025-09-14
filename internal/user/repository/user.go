package repositories

import (
	"errors"

	"github.com/Minhajxdd/Ephemr/internal/user/model"
	"gorm.io/gorm"
)

// UserRepository defines the methods for interacting with user table in the database.
type UserRepository interface {
	// Create a new user record in the database.
	Create(user *model.User) error

	// FindByEmail retrieves a user record based on their email address.
	FindByEmail(email string) (*model.User, error)

	// GetByID retrieves a user record by their unique ID.
	GetByID(id uint) (*model.User, error)

	// SetVerifyStatus updates the verification status of a user's record.
	SetVerifyStatus(id uint, status bool) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(u *model.User) error {
	return r.db.Create(u).Error
}

func (r *userRepo) FindByEmail(email string) (*model.User, error) {
	var u model.User
	if err := r.db.Where(&model.User{Email: email}).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) GetByID(id uint) (*model.User, error) {
	var u model.User
	if err := r.db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) SetVerifyStatus(id uint, status bool) (*model.User, error) {
	var user model.User

	result := r.db.Model(&model.User{}).
		Where("id = ?", id).
		Update("is_verified", status)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
