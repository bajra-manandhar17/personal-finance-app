package user

import (
	"context"
	"fmt"
	"log"

	"github.com/bajra-manandhar17/personal-finance-app/internal/db/model"
	"github.com/bajra-manandhar17/personal-finance-app/internal/excep"
	"github.com/bajra-manandhar17/personal-finance-app/internal/repository/userrepo"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepo *userrepo.UserRepo
}

func (u *UserServiceImpl) RegisterNewUser(
	ctx context.Context,
	req RegisterNewUserReq,
) error {
	if u.userRepo == nil {
		log.Fatal("queries is not initialized in UserServiceImpl")
	}

	emailCount, err := (*u.userRepo).DoesEmailExist(ctx, req.Email)
	if err != nil {
		return fmt.Errorf("error checking if email exists: %w", err)
	}

	if emailCount {
		return excep.DomainExcep{
			Type:    excep.EXCEP_EMAIL_EXISTS,
			Details: "email already exists",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
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

	if err := (*u.userRepo).CreateUser(ctx, &user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (u *UserServiceImpl) GetUser(
	ctx context.Context,
	userId string,
) (*model.Users, error) {
	if u.userRepo == nil {
		log.Fatal("userRepo is not initialized in UserServiceImpl")
	}

	userExists, err := (*u.userRepo).DoesUserExist(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("error checking if user exists: %w", err)
	}

	if !userExists {
		return nil, excep.DomainExcep{
			Type:    excep.EXCEP_USER_NOT_FOUND,
			Details: "user not found",
		}
	}

	userInfo, err := (*u.userRepo).GetUser(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return userInfo, nil
}
