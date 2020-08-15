package users

import (
	"fmt"
	"strings"

	"github.com/hsimao/bookstore_users-api/datasources/mysql/users_db"
	"github.com/hsimao/bookstore_users-api/utils/dateUtils"
	"github.com/hsimao/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]

	if result != nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dateUtils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {

		// handle mysql email 重複
		if strings.Contains(saveErr.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exists", user.Email))
		}

		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", saveErr.Error()),
		)
	}

	userID, idErr := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", idErr.Error()),
		)
	}

	user.Id = userID
	return nil
}
