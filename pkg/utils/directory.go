package utils

import (
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// Return all files and dirs in [dirPath], optional matching suffix
func ListDir(dirPth string, suffix string) (dirs []string, files []string, err error) {
	files = make([]string, 0, 10)
	dirs = make([]string, 0, 10)
	dir, err := os.ReadDir(dirPth)
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

// Get all files in the specified directory and all subdirectories, matching suffix filtering.
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

	_dir, err := os.ReadDir(dir)
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

// Copy directory deeply
func CopyDir(srcPath, desPath string) error {
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return ErrSrcPathInvalid
		}
	}
	if desInfo, err := os.Stat(desPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return ErrDstPathInvalid
		}
	}
	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return ErrSrcSameAsDst
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
				return MakeDir(destNewPath, 0777)
			}
		}
		return nil
	})

	return err
}

func MakeDir(dir string, perm fs.FileMode) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, perm); err != nil { //os.ModePerm
			Error(1, "MakeDir failed: %v", err)
			return err
		}
	}
	return nil
}

func RunInDir(dir string, cmd *exec.Cmd) (err error) {
	os.Chdir(dir)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func DeleteThingsInDir(targetDir string) error {
	dir, err := os.ReadDir(targetDir)
	if err != nil {
		return err
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{targetDir, d.Name()}...))
	}
	return err
}

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}
