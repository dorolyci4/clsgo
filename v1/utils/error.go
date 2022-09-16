package utils

import (
	"errors"
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
)