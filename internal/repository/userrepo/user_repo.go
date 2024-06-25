package userrepo

import (
	"context"

	"github.com/bajra-manandhar17/personal-finance-app/internal/db/model"
	"github.com/bajra-manandhar17/personal-finance-app/internal/db/query"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.Users) error
	GetUser(ctx context.Context, userId string) (*model.Users, error)
	DoesUserExist(ctx context.Context, userId string) (bool, error)
	DoesEmailExist(ctx context.Context, email string) (bool, error)
}

func NewUserRepo(query *query.Query) UserRepo {
	return &UserRepoImpl{
		query: query,
	}
}
