package errors

import "errors"

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

//func IsNotFoundError(err) bool {
//	return errors.Is()
//}

type UnknownError struct {
	Message string
}

func (e UnknownError) Error() string {
	return e.Message
}

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
