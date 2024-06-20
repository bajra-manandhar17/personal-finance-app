package userrepo

import (
	"context"

	"github.com/bajra-manandhar17/personal-finance-app/internal/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, userId string) (*model.User, error)
	DoesEmailExist(ctx context.Context, email string) (bool, error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}
