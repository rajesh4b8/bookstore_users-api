package users

import (
	"fmt"
	"time"

	"github.com/rajesh4b8/bookstore_users-api/datasources/mysql/users_db"
	dateutils "github.com/rajesh4b8/bookstore_users-api/utils/date_utils"
	"github.com/rajesh4b8/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email) VALUES($1, $2, $3) RETURNING user_id, date_created"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("Server error when prepare stmt")
	}
	// user.DateCreated = dateutils.GetNowString()
	var userId int64
	var dateCreated time.Time
	saveErr := stmt.QueryRow(user.FirstName, user.LastName, user.Email).Scan(&userId, &dateCreated)
	if saveErr != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("Error while inserting the user")
	}
	// current := usersDB[user.Id]
	// if current != nil {
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	// }

	// usersDB[user.Id] = user
	user.Id = userId
	user.DateCreated = dateCreated.Format(dateutils.ApiDateLayout)
	return nil
}
