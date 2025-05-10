package repository

import (
	"context"
	"doctor-on-demand/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
type userRepo struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepo{db: db}
}

// CreateUser saves a new user to the database
func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByEmail fetches a user by email from the database
func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
