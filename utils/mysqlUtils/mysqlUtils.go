package mysqlUtils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/hsimao/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing mysql response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewInternalServerError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
