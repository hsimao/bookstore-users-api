package services

import (
	"github.com/hsimao/bookstore_users-api/domain/users"
	"github.com/hsimao/bookstore_users-api/utils/cryptoUtils"
	"github.com/hsimao/bookstore_users-api/utils/dateUtils"
	"github.com/hsimao/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = dateUtils.GetNowDBFormat()
	user.Password = cryptoUtils.GetMd5(user.Password)
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email

	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
