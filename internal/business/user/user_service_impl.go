package user

import (
	"context"
	"fmt"
	"log"

	"github.com/bajra-manandhar17/personal-finance-app/internal/db/model"
	"github.com/bajra-manandhar17/personal-finance-app/internal/db/query"
	"github.com/bajra-manandhar17/personal-finance-app/internal/excep"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	query *query.Query
}

func (u *UserServiceImpl) RegisterNewUser(ctx context.Context, req RegisterNewUserReq) error {
	if u.query == nil {
		log.Fatal("queries is not initialized in UserServiceImpl")
	}

	userEmail := u.query.Users.Email

	emailCount, err := u.newUserQuery(ctx).Where(userEmail.Eq(req.Email)).Count()
	if err != nil {
		return fmt.Errorf("error checking if email exists: %w", err)
	}

	if emailCount > 0 {
		return excep.DomainExcep{
			Type:    excep.EXCEP_EMAIL_EXISTS,
			Details: "Email already exists",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	userId, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("error generating user id: %w", err)
	}

	user := model.Users{
		UserID:         userId.String(),
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
	}

	if err := u.newUserQuery(ctx).Create(&user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (u *UserServiceImpl) GetUser(ctx context.Context, userId string) (*model.Users, error) {
	if u.query == nil {
		log.Fatal("userRepo is not initialized in UserServiceImpl")
	}

	userID := u.query.Users.UserID

	userInfo, err := u.newUserQuery(ctx).Where(userID.Eq(userId)).First()
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return userInfo, nil
}

func (u *UserServiceImpl) newUserQuery(ctx context.Context) query.IUsersDo {
	return u.query.Users.WithContext(ctx)
}
