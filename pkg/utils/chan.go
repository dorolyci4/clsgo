package utils

import (
	"context"
	"time"
)

// General Chan write function
func WriteChanWithTimeout[T any](ctx context.Context, c chan T, data T, timeouts ...time.Duration) error {
	ms := time.Microsecond * 100
	if len(timeouts) > 0 && timeouts[0] > 0 {
		ms = timeouts[0]
	}
	timeout := time.NewTimer(time.Microsecond * time.Duration(ms))
	select {
	case <-ctx.Done():
		// ... exit
		return ErrChanOptCanceled
	case c <- data:
		return nil
	case <-timeout.C:
		return ErrChanWriteTimeout
	}
}

func ReadChanWithTimeout[T any](ctx context.Context, c <-chan T, timeouts ...time.Duration) (T, error) {
	ms := time.Microsecond * 100
	if len(timeouts) > 0 && timeouts[0] > 0 {
		ms = timeouts[0]
	}
	var empty T
	timeout := time.NewTimer(time.Microsecond * time.Duration(ms))
	select {
	case <-ctx.Done():
		// ... exit
		return empty, ErrChanOptCanceled
	case data := <-c:
		return data, nil
	case <-timeout.C:
		return empty, ErrChanReadTimeout
	}
}
