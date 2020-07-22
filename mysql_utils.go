package utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return NewNotFoundError("no record matching given id")
		}
		return NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return NewBadRequestError(sqlErr.Message)
	}

	return NewInternalServerError("error processing request")
}
