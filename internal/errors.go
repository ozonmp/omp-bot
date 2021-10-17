package internal

import "errors"

type UserError struct {
	Err error
}

func NewUserError(err string) UserError {
	return UserError{
		Err: errors.New(err),
	}
}

func (e UserError) Error() string {
	return e.Err.Error()
}
