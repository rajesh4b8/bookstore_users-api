package services

import (
	"github.com/rajesh4b8/bookstore_users-api/domain/users"
	"github.com/rajesh4b8/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := users.User{
		Id: userId,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}
