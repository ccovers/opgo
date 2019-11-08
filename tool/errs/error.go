package errs

import (
	"fmt"
)

type Error struct {
	err           error // original error
	stack         string
	code, message string
}

func New(code, message string) *Error {
	return &Error{code: code, message: message}
}

func Trace(err error) *Error {
	if e, ok := err.(*Error); ok {
		if e.Stack() == "" {
			e.stack = Stack(3)
		}
		return e
	} else {
		return &Error{
			err:   err,
			stack: Stack(3),
		}
	}
}

func Tracef(format string, args ...interface{}) *Error {
	return &Error{
		err:   fmt.Errorf(format, args...),
		stack: Stack(3),
	}
}

func (err *Error) Error() string {
	if err.err != nil {
		return err.err.Error()
	} else {
		return err.code + `: ` + err.message
	}
}

func (err *Error) Stack() string {
	return err.stack
}

func (err *Error) Trace() *Error {
	err.stack = Stack(3)
	return err
}

func (err *Error) Code() string {
	return err.code
}

func (err *Error) Message() string {
	return err.message
}

func (err *Error) SetCodeMessage(code, message string) *Error {
	err.code, err.message = code, message
	return err
}

func (err *Error) GetError() error {
	return err.err
}

func (err *Error) SetError(e error) *Error {
	err.err = e
	return err
}
