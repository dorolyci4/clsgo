package utils_test

import (
	"os"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/utils"
)

func Test_exit(t *testing.T) {
	ExampleSetupExitNotify()
}

func ExampleSetupExitNotify() {
	gtest.C(&testing.T{}, func(t *gtest.T) {
		exit := make(chan os.Signal, 1)
		utils.SetupExitNotify(exit)
		go func(x chan os.Signal) {
			time.Sleep(time.Second * 2)
			x <- os.Interrupt
		}(exit)
		e := <-exit
		t.Assert(e, os.Interrupt)
	})
}
