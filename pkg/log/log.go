// Package log provides glog functions.
package log

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"path"
	"runtime"
	"strconv"
)

func reverse(a []any) []any {
	prefix := a[len(a)-1]
	for i := len(a) - 1; i >= 1; i-- {
		a[i] = a[i-1]
	}
	a[0] = prefix
	return a
}

func Info(v ...any) {
	var ctx = context.TODO()

	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log().Info(ctx, v...)
}

func Infof(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Infof(ctx, s+f, v...)
}

func Debug(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log().Debug(ctx, v...)
}

func Debugf(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Debugf(ctx, s+f, v...)
}
func Warning(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log().Warning(ctx, v...)
}

func Warningf(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Warningf(ctx, s+f, v...)
}

func Error(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log().Error(ctx, v...)
}

func Errorf(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Errorf(ctx, s+f, v...)
}

func Panic(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log().Panic(ctx, v...)
}

func Panicf(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Panicf(ctx, s+f, v...)
}

func Fatal(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log().Fatal(ctx, v...)
}

func Fatalf(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Fatalf(ctx, s+f, v...)
}
