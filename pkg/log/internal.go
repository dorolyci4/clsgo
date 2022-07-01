/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-15 17:31:09
 * @LastEditTime    : 2022-07-01 15:42:58
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/log/internal.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package log

import (
	"fmt"
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

func Info(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("%s[INFO][%v:%v %v] %s%s\n", ANSI_GREEN, path.Base(file), line, path.Base(name), fmt.Sprintf(format, args...), ANSI_RESET)
}

func Warning(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("%s[WARN][%v:%v %v] %s%s\n", ANSI_MAGENTA, path.Base(file), line, path.Base(name), fmt.Sprintf(format, args...), ANSI_RESET)
}

func Error(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("%s[ERRO][%v:%v %v] %s%s\n", ANSI_RED, path.Base(file), line, path.Base(name), fmt.Sprintf(format, args...), ANSI_RESET)
}

func Important(format string, args ...interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Printf("%s[IMPO][%v:%v %v] %s%s\n", ANSI_BLUE, path.Base(file), line, path.Base(name), fmt.Sprintf(format, args...), ANSI_RESET)
}
