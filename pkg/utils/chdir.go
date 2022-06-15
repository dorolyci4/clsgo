package utils

import (
	"os"
	"path/filepath"
)

func ChdirToPos() (err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}
	err = os.Chdir(dir)
	return
}
