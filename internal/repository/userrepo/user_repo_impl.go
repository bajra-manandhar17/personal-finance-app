package userrepo

import (
	"context"
	"fmt"

	"github.com/bajra-manandhar17/personal-finance-app/internal/model"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func (u *UserRepoImpl) CreateUser(ctx context.Context, user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (u *UserRepoImpl) GetUser(ctx context.Context, userId string) (*model.User, error) {
	var user model.User

	if err := u.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return &user, nil
}

func (u *UserRepoImpl) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := u.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, fmt.Errorf("error checking if email exists: %w", err)
	}

	return count > 0, nil
}
