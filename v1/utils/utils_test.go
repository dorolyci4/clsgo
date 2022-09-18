package utils_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/utils"
)

func init() {
	log.Green("Runing test cases of charset")
}
func clean() {
	utils.DeletePath("logs")
	utils.DeleteFiles(utils.Cwd(), "/*.yaml$")
	utils.DeleteFiles(utils.Cwd(), "/*.xml$")
}

func TestDirectoryCases(t *testing.T) {
	log.Green("Running dir test cases")
	dir := "test"
	gtest.C(t, func(t *gtest.T) {
		t.Run("MakeDir", func(to *testing.T) {
			t.Assert(utils.DeleteThingsInDir(dir), nil)
			t.Assert(utils.MakeDir(dir, 0755), nil)
			t.Assert(utils.MakeDir(dir, 0755), nil)
		})
		t.Run("CopyFile", func(to *testing.T) {
			n1, err := utils.CopyFile("check.go", filepath.Join(dir, "check.go"))
			t.Assert(err, nil)
			n2, err := utils.CopyFile("check.go", filepath.Join(dir, "check.go"))
			t.Assert(err, nil)
			t.Assert(n1, n2)
		})
		t.Run("ListDir", func(to *testing.T) {
			dirs, files, err := utils.ListDir(dir, ".go")
			t.Assert(err, nil)
			t.Assert(len(dirs), 0)
			t.Assert(len(files), 1)
		})
		t.Run("CreateFile", func(to *testing.T) {
			utils.MakeDir(filepath.Join(dir, "from/1/2/3/4/5"), 0755)
			t.Assert(utils.CreateFile(filepath.Join(dir, "from/1/1.txt"), "1", 0755), nil)
			t.Assert(utils.CreateFile(filepath.Join(dir, "from/1/2/2.md"), "2", 0755), nil)
			t.Assert(utils.CreateFile(filepath.Join(dir, "from/1/2/3/3.c"), "3", 0755), nil)
			t.Assert(utils.CreateFile(filepath.Join(dir, "from/1/2/3/4/4.java"), "4", 0755), nil)
			t.Assert(utils.CreateFile(filepath.Join(dir, "from/1/2/3/4/5/5.html"), "5", 0755), nil)
		})
		t.Run("WalkDir", func(to *testing.T) {
			utils.MakeDir(filepath.Join(dir, "subdir"), 0755)

			dirs, files, err := utils.WalkDir(dir, ".go", ".txt")

			t.Assert(err, nil)
			t.Assert(len(files), 2)
			t.Assert(len(dirs), 8)

		})
		t.Run("CopyDir", func(to *testing.T) {
			t.Assert(utils.CopyDir(dir+"/from", dir+"/subdir"), nil)
		})
		t.Run("CopyToNewDir", func(to *testing.T) {
			t.Assert(utils.CopyToNewDir(dir+"/from", dir+"/to"), nil)
		})
		t.Run("RunIn", func(to *testing.T) {
			c := utils.Cwd()
			t.Assert(utils.RunIn(dir+"/from", "go", "version"), nil)
			utils.ChdirToPos(c)
		})
		t.Run("DeletePath", func(to *testing.T) {
			t.Assert(utils.DeletePath(dir), nil)
		})
	})
	gtest.C(t, func(t *gtest.T) {
	})
}

// Test result check tool:
// https://www.qqxiuzi.cn/bianma/zifuji.php

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

func TestChan(t *testing.T) {
	log.Green("Running channel test cases")
	channel := make(chan int)
	ctx := context.TODO()
	gtest.C(t, func(t *gtest.T) {
		t.Run("Write", func(tcase *testing.T) {
			t.Assert(utils.WriteChanWithTimeout(ctx, channel, 3), utils.ErrChanWriteTimeout)
		})
		t.Run("Read", func(tcase *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)
			go func(ctx context.Context, wg *sync.WaitGroup) {
				x, err := utils.ReadChanWithTimeout(ctx, channel, time.Hour)
				t.Assert(err, nil)
				t.Assert(x, 3)
				wg.Done()
			}(ctx, &wg)
			t.Assert(utils.WriteChanWithTimeout(ctx, channel, 3), nil)
			wg.Wait()
		})
	})
}

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

func TestFileIsExisted(t *testing.T) {
	log.Green("Running file test cases")
	gtest.C(t, func(t *gtest.T) {
		log.Green(utils.Cwd())
		t.Assert(utils.FileIsExisted("./file.go"), true)
		t.Assert(utils.FileIsExisted("file.go"), true)
		t.Assert(utils.FileIsExisted("-This-Is-Not-Existed.go"), false)
		t.Assert(utils.FileIsExisted("../config"), true)
		t.Assert(utils.FileIsExisted("../ThisIsNotExisted"), false)
	})

	clean()
}
