package utils_test

import (
	"fmt"
	"testing"

	"github.com/lovelacelee/clsgo/pkg/utils"
)

func Test(t *testing.T) {
	ExampleSetupExitNotify()
}

func ExampleSetupExitNotify() {
	exit := make(chan bool, 1)
	utils.SetupExitNotify(exit)
	e := <-exit
	fmt.Println(e)
}
