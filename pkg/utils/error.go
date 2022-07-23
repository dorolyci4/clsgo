package utils

import (
	"errors"
)

var (
	ErrEOF           = errors.New("EOF")
	ErrUnexpectedEOF = errors.New("unexpected EOF")
	ErrNoProgress    = errors.New("multiple Read calls return no data or error")
)
