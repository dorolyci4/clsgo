package internal

import (
	"errors"
)

var (
	ErrChanWriteTimeout = errors.New("write chan time out")
)
