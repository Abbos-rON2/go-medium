package errs

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPassword  = New("invalid password")
	ErrInvalidToken     = New("invalid token")
	ErrUnauthorized     = New("unauthorized")
	ErrNotFound         = New("entity not found")
	ErrInternal         = New("internal error")
	ErrValidation       = New("validation error")
	ErrDatabaseError    = New("database error")
	ErrLikeAleadyExists = New("like already exists")
)

type Error struct {
	err error
	msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.err.Error(), e.msg)
}

func (e *Error) Unwrap() error {
	return e.err
}

func Errf(err error, tmpl error, args ...interface{}) error {
	return &Error{err: err, msg: fmt.Sprintf(tmpl.Error(), args...)}
}

func New(text string) *Error {
	return &Error{err: errors.New(text)}
}
