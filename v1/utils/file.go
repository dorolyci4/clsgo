package utils

import (
	"io/fs"
	"os"
)

func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	fi, _ := srcFile.Stat()
	perm := fi.Mode()
	srcFile.Close()

	input, err := os.ReadFile(src)
	if err != nil {
		return 0, err
	}

	err = os.WriteFile(des, input, perm)
	if err != nil {
		return 0, err
	}

	return int64(len(input)), nil
}

func MoveFile(src, des string) {
	if FileIsExisted(des) {
		os.Remove(des)
	}
	_, e := CopyFile(src, des)
	if e == nil {
		os.Remove(src)
	}
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func CreateFile(path string, data string, mode fs.FileMode) error {
	MakeDir(path, 0755, true)
	return os.WriteFile(path, []byte(data), mode)
}
