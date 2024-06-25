package user

import (
	"context"

	"github.com/bajra-manandhar17/personal-finance-app/internal/db/model"
	"github.com/bajra-manandhar17/personal-finance-app/internal/repository/userrepo"
)

type UserService interface {
	RegisterNewUser(ctx context.Context, req RegisterNewUserReq) error
	GetUser(ctx context.Context, userId string) (*model.Users, error)
}

type RegisterNewUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserServiceOpts struct {
	UserRepo *userrepo.UserRepo
}

func NewUserService(opts *UserServiceOpts) UserService {
	if opts == nil {
		return &UserServiceImpl{}
	}

	return &UserServiceImpl{
		userRepo: opts.UserRepo,
	}
}
