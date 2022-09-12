package utils

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// Get current project's absolute path
func GetCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// Unix-like system temp environment
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// Chdir to the real application exist path
func ChdirToPos() (err error) {
	err = os.Chdir(GetCurrentAbPath())
	return
}

func PathFix(p string) string {
	return strings.ReplaceAll(p, "\\", "/")
}

func PatchClean(p string) string {
	return path.Clean(PathFix(p))
}

func PathJoin(elem ...string) string {
	for i, e := range elem {
		if e != "" {
			return path.Clean(PathFix(strings.Join(elem[i:], "/")))
		}
	}
	return ""
}

func PathReplace(path string, from string, to string) string {
	return strings.Replace(PathFix(path), from, to, 1)
}

func HomePath() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func TempPath(p string) string {
	path := filepath.Join(os.TempDir(), p)
	MakeDir(path, 0755)
	return path
}

func Cwd() string {
	cwd, _ := os.Getwd()
	return cwd
}
