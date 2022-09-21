package utils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func sufixes(suffix ...string) []string {
	result := make([]string, 0)
	if !IsEmpty(suffix) {
		for _, s := range suffix {
			//Ignore case for suffix matches
			result = append(result, strings.ToUpper(s))
		}
		return result
	}
	return nil
}

func inSuffixes(s string, slist []string) bool {
	for _, suf := range slist {
		if strings.HasSuffix(strings.ToUpper(s), suf) { //file matches
			return true
		}
	}
	return false
}

// Return all files and dirs in [dirPath], optional matching suffix, sub directories ignored
func ListDir(dirPth string, suffix ...string) (dirs []string, files []string, err error) {
	dirPth = PathFix(dirPth)
	files = make([]string, 0, 10)
	dirs = make([]string, 0, 10)
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}
	PthSep := string(os.PathSeparator)
	sufList := sufixes(suffix...)
	for _, fi := range dir {
		if fi.IsDir() { // directory
			dirs = append(dirs, dirPth+PthSep+fi.Name())
		}
		if !IsEmpty(sufList) { //file matches
			if inSuffixes(fi.Name(), sufList) {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return dirs, files, nil
}

// Get all files in the specified directory and all subdirectories, matching suffix filtering.
func WalkDir(dirPth string, suffix ...string) (dirs []string, files []string, err error) {
	dirPth = PathFix(dirPth)
	files = make([]string, 0)
	dirs = make([]string, 0)
	suffixes := sufixes(suffix...)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			dirs = append(dirs, filename)
		}
		if !IsEmpty(suffix) {
			if inSuffixes(fi.Name(), suffixes) {
				files = append(files, filename)
			}
		} else {
			files = append(files, filename)
		}
		return err
	})
	return dirs, files, err
}

// Only the contents of the source directory are recursively copied to the destination path，
// source directory itself not copied. Returns error if source/destination path does not exist.
// And srcPath must be different from desPath.
func CopyDir(srcPath, desPath string) error {
	srcPath = PathFix(srcPath)
	desPath = PathFix(desPath)
	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return ErrSrcSameAsDst
	}
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
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if path == srcPath {
			return nil
		}
		destNewPath := PathReplace(path, srcPath, desPath)
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

// desPath is not exist
func CopyToNewDir(srcPath, desPath string) error {
	srcPath = PathFix(srcPath)
	desPath = PathFix(desPath)
	target := filepath.Join(desPath, filepath.Base(srcPath))
	if !FileIsExisted(target) {
		MakeDir(target, 0777)
	}
	return CopyDir(srcPath, target)
}

// Create directory: The path parameter can be a directory or a file.
// If it is a file path, it will be truncated by the Dir function
func MakeDir(path string, perm fs.FileMode, isfile ...bool) error {
	var dir string
	if Param(isfile, false) {
		dir = PathFix(filepath.Dir(path))
	} else {
		dir = PathFix(path)
	}
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, perm); err != nil { //os.ModePerm
			// Error(1, "err: %v", err)
			return err
		}
	}
	return nil
}

// cmd := exec.Command("git", "push", remote, "--tags", "--force")
// utils.RunInDir(w.Filesystem.Root(), cmd)
func RunInDir(dir string, cmd *exec.Cmd) (err error) {
	dir = PathFix(dir)
	os.Chdir(dir)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func RunIn(dir string, c string, args ...string) error {
	dir = PathFix(dir)
	cmd := exec.Command(c, args...)
	return RunInDir(dir, cmd)
}

func Exec(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return err
	}
	// stdout
	logScan := bufio.NewScanner(stdout)
	go func() {
		for logScan.Scan() {
			fmt.Println(logScan.Text())
		}
	}()
	// stderr
	errBuf := bytes.NewBufferString("")
	scan := bufio.NewScanner(stderr)
	for scan.Scan() {
		s := scan.Text()
		errBuf.WriteString(s)
		errBuf.WriteString("\n")
	}

	cmd.Wait()
	if !cmd.ProcessState.Success() {
		return errors.New(errBuf.String())
	}
	return nil
}

// Delete all files in the directory, left an empty directory(targetDir)
func DeleteThingsInDir(targetDir string) error {
	targetDir = PathFix(targetDir)
	dir, err := os.ReadDir(targetDir)
	if err == nil {
		for _, d := range dir {
			os.RemoveAll(path.Join([]string{targetDir, d.Name()}...))
		}
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// Delete the directory
func DeletePath(path string) error {
	path = PathFix(path)
	if FileIsExisted(path) && IsDir(path) {
		DeleteThingsInDir(path)
	}
	return os.RemoveAll(path)
}

func DeleteFiles(path string, regexpr string) {
	_, files, _ := ListDir(path)
	reg, err := regexp.Compile(regexpr)
	if err != nil {
		return
	}
	for _, f := range files {
		name := filepath.Base(f)
		if reg.MatchString(name) {
			fp := PathFix(f)
			os.Remove(fp)
		}
	}
}

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}
