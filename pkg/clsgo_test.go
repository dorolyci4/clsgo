// Test package for clsgo
package clsgo_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
	"github.com/lovelacelee/clsgo/pkg/version"
)

func TestClsgo(t *testing.T) {
	v := version.Version
	log.Green("Running clsgo test cases")
	gtest.C(t, func(t *gtest.T) {
		t.Assert(v, utils.NumberToString(version.Major)+"."+utils.NumberToString(version.Minor)+"."+utils.NumberToString(version.Patch))
	})
}
