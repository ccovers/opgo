package errs

import (
	"fmt"
	"time"
)

type Error struct {
	err      error  `json:"err"` // original error
	stack    string `json:"stack"`
	code     string `json:"code"`
	message  string `json:"message"`
	time     string `json:"time"`
	filename string `json:"filename"`
	line     int    `json:"line"`
	funcname string `json:"funcname"`
}

func New(code, message string) *Error {
	fileName, line, functionName := GetFileInfo()
	return &Error{
		code:     code,
		message:  message,
		time:     time.Now().Format("2006-01-02 15:04:05"),
		filename: fileName,
		line:     line,
		funcname: functionName,
	}
}

func Trace(err error) *Error {
	if e, ok := err.(*Error); ok {
		if e.Stack() == "" {
			e.stack = Stack(3)
		}
		return e
	} else {
		fileName, line, functionName := GetFileInfo()
		return &Error{
			err:      err,
			stack:    Stack(3),
			time:     time.Now().Format("2006-01-02 15:04:05"),
			filename: fileName,
			line:     line,
			funcname: functionName,
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
		return fmt.Sprintf(`%s:%s [%s:%d:%s]`, err.code, err.message,
			err.filename, err.line, err.funcname)
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
