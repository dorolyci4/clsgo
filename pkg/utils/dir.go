package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

//Return all files and dirs in [dirPath], optional matching suffix
func ListDir(dirPth string, suffix string) (dirs []string, files []string, err error) {
	files = make([]string, 0, 10)
	dirs = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}
	PthSep := string(os.PathSeparator)
	//Ignore case for suffix matches
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() { // directory
			dirs = append(dirs, dirPth+PthSep+fi.Name())
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //file matches
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return dirs, files, nil
}

//Get all files in the specified directory and all subdirectories, matching suffix filtering.
func WalkDir(dirPth, suffix string) (dirs []string, files []string, err error) {
	files = make([]string, 0, 30)
	dirs = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //Ignore case for suffix matches
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			dirs = append(dirs, filename)
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return err
	})
	return dirs, files, err
}

// Get all files under the specified path, search only the current path,
// do not enter the next level of directory, support suffix filtering (suffix
// is empty, not filtered)
func ListDirFiles(dir, suffix string) (files []string, err error) {
	files = []string{}

	_dir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToLower(suffix) //suffix filtering

	for _, _file := range _dir {
		if _file.IsDir() {
			continue //
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(_file.Name()), suffix) {
			//
			files = append(files, path.Join(dir, _file.Name()))
		}
	}

	return files, nil
}

func CopyDir(srcPath, desPath string) error {

	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("srcPath is not a valid directory path")
		}
	}

	if desInfo, err := os.Stat(desPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return errors.New("destPath is not a valid directory path")
		}
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return errors.New("destPath must be different from srcPath")
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if path == srcPath {
			return nil
		}

		destNewPath := strings.Replace(path, srcPath, desPath, -1)

		if !f.IsDir() {
			CopyFile(path, destNewPath)
		} else {
			if !FileIsExisted(destNewPath) {
				return MakeDir(destNewPath)
			}
		}

		return nil
	})

	return err
}

func CopyFile(src, des string) (written int64, err error) {

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	fi, _ := srcFile.Stat()
	perm := fi.Mode()
	srcFile.Close()

	input, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(des, input, perm)
	if err != nil {
		return 0, err
	}

	return int64(len(input)), nil
}

func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}

func RunInDir(dir string, cmd *exec.Cmd) (err error) {
	os.Chdir(dir)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func DeleteThingsInDir(targetDir string) error {
	dir, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return err
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{targetDir, d.Name()}...))
	}
	return err
}
