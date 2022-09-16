package utils_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/utils"
)

func TestChan(t *testing.T) {
	log.Green("Running channel test cases")
	channel := make(chan int)
	ctx := context.TODO()
	gtest.C(t, func(t *gtest.T) {
		t.Run("Write", func(tcase *testing.T) {
			t.Assert(utils.WriteChanWithTimeout(ctx, channel, 3), utils.ErrChanWriteTimeout)
		})
		t.Run("Read", func(tcase *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)
			go func(ctx context.Context, wg *sync.WaitGroup) {
				x, err := utils.ReadChanWithTimeout(ctx, channel, time.Hour)
				t.Assert(err, nil)
				t.Assert(x, 3)
				wg.Done()
			}(ctx, &wg)
			t.Assert(utils.WriteChanWithTimeout(ctx, channel, 3), nil)
			wg.Wait()
		})
	})
}
