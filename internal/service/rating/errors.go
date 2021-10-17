package rating

import "errors"

//UserError - если такой вариант подходит, то надо выносить на глобальный уровень как базовые типы ошибок
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
