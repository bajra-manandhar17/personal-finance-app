package userrepo

import (
	"context"
	"fmt"

	"github.com/bajra-manandhar17/personal-finance-app/internal/db/model"
	"github.com/bajra-manandhar17/personal-finance-app/internal/db/query"
)

type UserRepoImpl struct {
	query *query.Query
}

func (u *UserRepoImpl) CreateUser(ctx context.Context, user *model.Users) error {
	if err := u.newUserQuery(ctx).Create(user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (u *UserRepoImpl) GetUser(ctx context.Context, userId string) (*model.Users, error) {
	user, err := u.newUserQuery(ctx).Where(u.query.Users.UserID.Eq(userId)).First()
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}

func (u *UserRepoImpl) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	count, err := u.newUserQuery(ctx).Where(u.query.Users.Email.Eq(email)).Count()
	if err != nil {
		return false, fmt.Errorf("error checking if email exists: %w", err)
	}

	return count > 0, nil
}

func (u *UserRepoImpl) newUserQuery(ctx context.Context) query.IUsersDo {
	return u.query.Users.WithContext(ctx)
}
