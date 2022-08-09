package utils_test

import (
	"errors"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
	"testing"
)

var errTest = errors.New("some error")

func Test_check(t *testing.T) {
	ExampleWarnIfError()
	Example()
}

func ExampleWarnIfError() {
	utils.WarnIfError(errTest)
	utils.WarnIfError(errTest, log.Warningf, "%s %s", "warning", "message")
}

func Example() {
	utils.InfoIfError(errTest, log.Infof, "%s %s", "warning", "message")
	utils.ImptIfError(errTest, log.Infof, "%s %s", "warning", "message")
	// utils.ExitIfError(errTest, log.Infof, "%s %s", "error", "message")
}
