package utils

import (
	"os"
	"path"
	"runtime"
	"strconv"
)

func checkerWithoutHeader(err error, fn ...any) {
	if len(fn) == 1 {
		fn[0].(Loggerf)("%s", err)
		return
	} else if len(fn) == 2 {
		fn[0].(Loggerf)(fn[1].(string)+"%s", err)
		return
	}
	switch fn[0].(type) {
	case Loggerf:
		fn[0].(Loggerf)(fn[1].(string), fn[2:]...)
	}
}

func checker(err error, fn ...any) {
	pc, file, line, _ := runtime.Caller(2)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	if len(fn) == 1 {
		fn[0].(Loggerf)(s+"%s", err)
		return
	} else if len(fn) == 2 {
		fn[0].(Loggerf)(s+fn[1].(string)+"%s", err)
		return
	}
	switch fn[0].(type) {
	case Loggerf:
		fn[0].(Loggerf)(s+fn[1].(string), fn[2:]...)
	}
}

// CheckIfError should be used to naively panics if an error is not nil. Eg:
// utils.ExitIfError(errTest, log.Warningf, "%s %s", "warning", "message")
func ExitIfError(err error, fn ...any) error {
	if err == nil {
		return nil
	}
	if len(fn) >= 1 {
		checker(err, fn...)
	} else {
		Error(2, "%s", err)
	}
	os.Exit(1)
	return os.ErrClosed
}

func ExitIfErrorWithoutHeader(err error, fn ...any) error {
	if err == nil {
		return nil
	}
	if len(fn) >= 1 {
		checkerWithoutHeader(err, fn...)
	} else {
		Error(2, "%s", err)
	}
	os.Exit(1)
	return os.ErrClosed
}

// Only output in terminal in [Error] message, Eg:
// utils.IfError(errTest, log.Errorf, "%s %s", "error", "error message")
func IfError(err error, fn ...any) error {
	if err == nil {
		return nil
	}
	if len(fn) >= 1 {
		checker(err, fn...)
	} else {
		Error(2, "%s", err)
	}
	return err
}

func IfErrorWithoutHeader(err error, fn ...any) error {
	if err == nil {
		return nil
	}
	if len(fn) >= 1 {
		checkerWithoutHeader(err, fn...)
	} else {
		Error(2, "%s", err)
	}
	return err
}

// Only output in terminal in [WARN] message, Eg:
// utils.WarnIfError(errTest, log.Warningf, "%s %s", "warning", "message")
func WarnIfError(err error, fn ...any) error {
	if err == nil {
		return nil
	}
	if len(fn) >= 1 {
		checker(err, fn...)
	} else {
		Warn(2, "%s", err)
	}
	return err
}

func WarnIfErrorWithoutHeader(err error, fn ...any) error {
	if err == nil {
		return nil
	}
	if len(fn) >= 1 {
		checkerWithoutHeader(err, fn...)
	} else {
		Warn(2, "%s", err)
	}
	return err
}

// Only output in terminal in [INFO] message, Eg:
// utils.InfoIfError(errTest, log.Infof, "%s %s", "warning", "message")
func InfoIfError(err error, fn ...any) error {
	if err == nil {
		return nil
	}
	if len(fn) >= 1 {
		checker(err, fn...)
	} else {
		Info(2, "%s", err)
	}
	return err
}

func InfoIfErrorWithoutHeader(err error, fn ...any) error {
	if err == nil {
		return nil
	}
	if len(fn) >= 1 {
		checkerWithoutHeader(err, fn...)
	} else {
		Info(2, "%s", err)
	}
	return err
}
