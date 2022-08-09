package utils

import (
	"os"
	"path"
	"runtime"
	"strconv"
)

// CheckIfError should be used to naively panics if an error is not nil.
func ExitIfError(err error, fn ...any) {
	if err == nil {
		return
	}
	if len(fn) >= 2 {
		checker(fn...)
	} else {
		Error("%s", err)
	}
	os.Exit(1)
}

func checker(fn ...any) {
	pc, file, line, _ := runtime.Caller(2)
	name := runtime.FuncForPC(pc).Name()
	s := "[CLSGO[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	switch fn[0].(type) {
	case Loggerf:
		fn[0].(Loggerf)(s+fn[1].(string), fn[2:]...)
	}
}

// Only output in terminal in [WARN] message
func WarnIfError(err error, fn ...any) {
	if err == nil {
		return
	}
	if len(fn) >= 2 {
		checker(fn...)
	} else {
		Warn("%s", err)
	}
}

// Only output in terminal in [INFO] message
func InfoIfError(err error, fn ...any) {
	if err == nil {
		return
	}
	if len(fn) >= 2 {
		checker(fn...)
	} else {
		Info("%s", err)
	}
}

// Only output in terminal in [IMPT] message
func ImptIfError(err error, fn ...any) {
	if err == nil {
		return
	}
	if len(fn) >= 2 {
		checker(fn...)
	} else {
		Impt("%s", err)
	}
}
