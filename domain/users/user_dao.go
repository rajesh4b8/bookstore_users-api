package users

import (
	"fmt"
	"strings"
	"time"

	"github.com/rajesh4b8/bookstore_users-api/datasources/mysql/users_db"
	dateutils "github.com/rajesh4b8/bookstore_users-api/utils/date_utils"
	"github.com/rajesh4b8/bookstore_users-api/utils/errors"
)

const (
	noRowsError     = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email) VALUES($1, $2, $3) RETURNING user_id, date_created"
	queryFetchUser  = "Select user_id, first_name, last_name, email, date_created from users where user_id = $1"
)

var (
// usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryFetchUser)
	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("Server error when prepare stmt")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("Error when fetching the user with id %d", user.Id))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("Server error when prepare stmt")
	}
	defer stmt.Close()

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
