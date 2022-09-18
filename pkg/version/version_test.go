package version_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/utils"
	"github.com/lovelacelee/clsgo/pkg/version"
)

func TestVersion(t *testing.T) {
	gtest.C(t, func(gt *gtest.T) {
		v := version.Version
		gt.Assert(v, utils.NumberToString(version.Major)+"."+
			utils.NumberToString(version.Minor)+"."+
			utils.NumberToString(version.Patch))
	})
}
