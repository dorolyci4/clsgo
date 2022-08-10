package internal

import (
	"context"
	"time"
)

// General Chan write function
func WriteChanWithTimeout[T any](ctx context.Context, c chan T, data T, timeouts ...uint) error {
	var ms uint = 100
	if len(timeouts) > 0 && timeouts[0] > 0 {
		ms = timeouts[0]
	}
	timeout := time.NewTimer(time.Microsecond * time.Duration(ms))
	select {
	case <-ctx.Done():
		// ... exit
		return nil
	case c <- data:
		return nil
	case <-timeout.C:
		return ErrChanWriteTimeout
	}
}
