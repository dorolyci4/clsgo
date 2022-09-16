package utils_test

import (
	"errors"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/utils"
	"testing"
)

var errMessage = "ok"
var errTest = errors.New(errMessage)

func Test_check(t *testing.T) {
	log.Green("Running check test cases")
	gtest.C(t, func(tc *gtest.T) {
		t.Run("Info", func(t *testing.T) {
			tc.Assert(utils.InfoIfErrorWithoutHeader(errTest, log.Green, "%s %s", "INFO", errMessage), errTest)
			tc.Assert(utils.InfoIfError(errTest, log.White, "%s %s", "INFO", errMessage), errTest)
		})
	})
}

func ExampleWarnIfError() {
	utils.WarnIfError(errTest)
	utils.WarnIfError(errTest, log.Warningf, "%s %s", "warning", "message")

}

func ExampleInfoIfError() {
	utils.InfoIfErrorWithoutHeader(errTest, log.Green, "%s %s", "INFO", "message")
	utils.InfoIfError(errTest, log.Infof, "%s %s", "INFO", "message")
}

func ExampleExitIfError() {
	utils.ExitIfError(errTest, log.Infof, "%s %s", "ERROR", "message")
}

func ExampleExitIfErrorWithoutHeader() {
	utils.ExitIfErrorWithoutHeader(errTest)
}
