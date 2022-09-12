package utils_test

import (
	"path/filepath"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
)

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
