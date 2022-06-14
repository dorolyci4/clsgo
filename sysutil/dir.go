package clsgo

import (
	"io/ioutil"
	"os"
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
		return nil
	})
	return dirs, files, err
}
