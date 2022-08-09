package utils_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/lovelacelee/clsgo/pkg/utils"
)

func Test_exit(t *testing.T) {
	ExampleSetupExitNotify()
}

func ExampleSetupExitNotify() {
	exit := make(chan os.Signal, 1)
	utils.SetupExitNotify(exit)
	go func(x chan os.Signal) {
		time.Sleep(time.Second * 2)
		x <- os.Interrupt
	}(exit)
	e := <-exit
	fmt.Println(e)
}
