// Package log provides [Infoi/Warni/Errori/Importanti] internal functions.
package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

var (
	ANSI_CYAN    = "\x1b[36;1m"
	ANSI_RESET   = "\x1b[0m"
	ANSI_DEFAULT = "\x1b[39;1m"
	ANSI_BLUE    = "\x1b[34;1m"
	ANSI_BLACK   = "\x1b[30;1m"
	ANSI_RED     = "\x1b[31;1m"
	ANSI_GREEN   = "\x1b[32;1m"
	ANSI_YELLOW  = "\x1b[33;1m"
	ANSI_WHITE   = "\x1b[37;1m"
	ANSI_MAGENTA = "\x1b[35;1m"
)

func isColored() bool {
	isColored := false

	switch force, ok := os.LookupEnv("CLICOLOR_FORCE"); {
	case ok && force != "0":
		isColored = true
	case ok && force == "0", os.Getenv("CLICOLOR") == "0":
		isColored = false
	}
	return isColored
}

func color(c string) string {
	if isColored() {
		return c
	} else {
		return ""
	}
}

func Infoi(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("%s[INFO][%v:%v %v] %s%s\n", color(ANSI_GREEN), path.Base(file), line, path.Base(name), fmt.Sprintf(format, args...), color(ANSI_RESET))
}

func Warni(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("%s[WARN][%v:%v %v] %s%s\n", color(ANSI_MAGENTA), path.Base(file), line, path.Base(name), fmt.Sprintf(format, args...), color(ANSI_RESET))
}

func Errori(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("%s[ERRO][%v:%v %v] %s%s\n", color(ANSI_RED), path.Base(file), line, path.Base(name), fmt.Sprintf(format, args...), color(ANSI_RESET))
}

func Importi(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("%s[IMPO][%v:%v %v] %s%s\n", color(ANSI_BLUE), path.Base(file), line, path.Base(name), fmt.Sprintf(format, args...), color(ANSI_RESET))
}
