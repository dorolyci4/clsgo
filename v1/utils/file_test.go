package utils_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/utils"
)

func TestFileIsExisted(t *testing.T) {
	log.Green("Running file test cases")
	gtest.C(t, func(t *gtest.T) {
		log.Green(utils.Cwd())
		t.Assert(utils.FileIsExisted("./file.go"), true)
		t.Assert(utils.FileIsExisted("file.go"), true)
		t.Assert(utils.FileIsExisted("-This-Is-Not-Existed.go"), false)
		t.Assert(utils.FileIsExisted("../version"), true)
		t.Assert(utils.FileIsExisted("../ThisIsNotExisted"), false)
	})
}
