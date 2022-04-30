package users

import (
	"github.com/rajesh4b8/bookstore_users-api/domain/users"
	"github.com/rajesh4b8/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

func GetUser() {

}

func FindUser() {

}
