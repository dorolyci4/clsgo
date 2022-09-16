package utils_test

import (
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/utils"
	"testing"
)

// Test result check tool:
// https://www.qqxiuzi.cn/bianma/zifuji.php

func init() {
	log.Green("Runing test cases of charset")
}

func TestStringConvert(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		chineseUtf8Str := "您好：Golang！"
		utf8Bytes := []byte{0xE6, 0x82, 0xA8, 0xE5, 0xA5, 0xBD, 0xEF, 0xBC, 0x9A, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0xef, 0xbc, 0x81}
		t.Assert(utf8Bytes, []byte(chineseUtf8Str))
		ansiBytes := []byte{0xC4, 0xFA, 0xBA, 0xC3, 0xA3, 0xBA, 0x47, 0x6F, 0x6C, 0x61, 0x6E, 0x67, 0xA3, 0xA1}
		gbkExpected, gb18030Expected := ansiBytes, ansiBytes
		gb2312Expected := "~{Dz:C#:~}Golang~{#!"

		t.Run("GB2312<->UTF8", func(_ *testing.T) {
			gb2312str, err := utils.StringConvert("UTF-8", "GB2312", chineseUtf8Str)
			t.Assert(err, nil)
			originstr, err := utils.StringConvert("GB2312", "UTF-8", gb2312str)
			t.Assert(err, nil)
			t.Assert(originstr, chineseUtf8Str)
			t.Assert(gb2312str, gb2312Expected)
		})
		t.Run("GBK<->UTF8", func(_ *testing.T) {
			gbkstr, err := utils.StringConvert("UTF-8", "GBK", chineseUtf8Str)
			t.Assert(err, nil)
			originstr, err := utils.StringConvert("GBK", "UTF-8", gbkstr)
			t.Assert(err, nil)
			t.Assert(originstr, chineseUtf8Str)
			t.Assert(utils.StringToHex(gbkstr), utils.BytesToHex(gbkExpected))
		})
		t.Run("GB18030<->UTF8", func(_ *testing.T) {
			gb18030str, err := utils.StringConvert("UTF-8", "GB18030", chineseUtf8Str)
			t.Assert(err, nil)
			originstr, err := utils.StringConvert("GB18030", "UTF-8", gb18030str)
			t.Assert(err, nil)
			t.Assert(originstr, chineseUtf8Str)
			t.Assert(utils.StringToHex(gb18030str), utils.BytesToHex(gb18030Expected))
		})
	})
}

func TestBytesDecode(t *testing.T) {
	utf8string := "안녕하세요" //go string encoded to utf-8
	utf8bytes := []byte(utf8string)

	gtest.C(t, func(t *gtest.T) {
		encbytes := utils.CharsetEncode(utils.UTF8, utf8bytes)
		outBytes, outStr, err := utils.BytesDecode(utils.UTF8, encbytes)
		t.Assert(err, nil)
		t.Assert(outStr, utf8string)
		t.Assert(outBytes, utf8bytes)
	})
}

func TestCoding(t *testing.T) {
	utf8string := "Здравствуйте,こんにちは" //go string encoded to utf-8
	utf8bytes := []byte(utf8string)
	gtest.C(t, func(t *gtest.T) {
		encbytes := utils.CharsetEncode(utils.UTF8, utf8bytes)
		decbytes := utils.CharsetDecode(utils.UTF8, encbytes)
		t.Assert(decbytes, utf8bytes)
	})
}

func TestDetermineDecode(t *testing.T) {
	utf8string := "Hello,您好，안녕하세요,Здравствуйте,こんにちは" //go string encoded to utf-8
	utf8bytes := []byte(utf8string)

	gtest.C(t, func(t *gtest.T) {
		outBytes, outStr, err := utils.DetermineDecode(utf8bytes)
		t.Assert(err, nil)
		t.Assert(outStr, utf8string)
		t.Assert(outBytes, utf8bytes)
	})
}
