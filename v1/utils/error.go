package utils

import (
	"errors"
	"fmt"
)

var (
	ErrEOF              = errors.New("EOF")
	ErrUnexpectedEOF    = errors.New("unexpected EOF")
	ErrNoProgress       = errors.New("multiple Read calls return no data or error")
	ErrSrcPathInvalid   = errors.New("srcPath is not a valid directory path")
	ErrDstPathInvalid   = errors.New("destPath is not a valid directory path")
	ErrSrcSameAsDst     = errors.New("destPath must be different from srcPath")
	ErrChanWriteTimeout = errors.New("write chan time out")
	ErrChanReadTimeout  = errors.New("read chan time out")
	ErrChanOptCanceled  = errors.New("chan operation canceled")
	ErrNotFound         = errors.New("404")
	ErrTcpDomain        = errors.New("domain server missing port in address")
)

type CError struct {
	message string
	err     error
}

func NewError(e error) *CError {
	errorWrapf := NewLayerFunctionErrorWrapf("Init", "NewError")
	return &CError{
		message: "",
		err:     errorWrapf(e, ""),
	}
}

func (e CError) Error() string {
	return e.message
}

func (e CError) Is(target error) bool {
	if target == nil || e.err == nil {
		return e.err == target
	}

	return errors.Is(e.err, target)
}

func (e *CError) Unwrap() error {
	u, ok := e.err.(interface {
		Unwrap() error
	})
	if !ok {
		return e.err
	}

	return u.Unwrap()
}

func (e *CError) Join(err error) *CError {
	if err == nil {
		return e
	}
	e.message += fmt.Sprintf("%v", err.Error())
	return e
}

func makeMessage(err error, layer, function, msg string) string {
	var message string
	var e CError
	if errors.As(err, &e) {
		message = fmt.Sprintf("[%s:%s] %s => %s", layer, function, msg, err.Error())
	} else {
		message = fmt.Sprintf("[%s:%s] %s => [Raw:Error] %v", layer, function, msg, err.Error())
	}

	return message
}

func Wrapf(err error, layer string, function string, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(format, args...)

	return CError{
		message: makeMessage(err, layer, function, msg),
		err:     err,
	}
}

type WrapfFuncWithLayerFunction func(err error, format string, args ...interface{}) error

// Usage: errorWrapf := utils.NewLayerFunctionErrorWrapf("Handler", "ListSubject")
func NewLayerFunctionErrorWrapf(layer string, function string) WrapfFuncWithLayerFunction {
	return func(err error, format string, args ...interface{}) error {
		return Wrapf(err, layer, function, format, args...)
	}
}
