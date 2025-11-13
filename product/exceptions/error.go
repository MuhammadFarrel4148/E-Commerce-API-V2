package exceptions

import "errors"

var ErrNotFound = errors.New("resource not found")

type error interface {
	Error() string
}

type ErrValidation struct {
	Details map[string]string
}

func (e *ErrValidation) Error() string {
	return "validation failed"
}
