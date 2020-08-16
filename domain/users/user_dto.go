package users

import (
	"strings"

	"github.com/hsimao/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json: "id"`
	FirstName   string `json: "firstName"`
	LastName    string `json: "lastName"`
	Email       string `json: "email"`
	DateCreated string `json: "dateCreated"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
