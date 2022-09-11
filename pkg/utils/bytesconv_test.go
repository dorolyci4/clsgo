package utils_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
)

func TestHex(t *testing.T) {
	log.Green("Runing bytesconv test cases")
	srcBytes := []byte{0x5B, 0x47, 0x4F, 0x5D}
	srcStr := "[GO]"
	srcHexStr := "5b474f5d"
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.BytesToHex(srcBytes), []byte(srcHexStr))
		t.Assert(utils.HexToBytes([]byte(srcHexStr)), srcBytes)

		hex := utils.StringToHex(srcStr)
		t.Assert(utils.HexToBytes(hex), srcBytes)
		str := utils.HexToString(hex)
		t.Assert(str, srcStr)

		t.Assert(utils.HexString([]byte(srcStr)), srcHexStr)
		utils.HexDump(srcBytes)
	})
}
