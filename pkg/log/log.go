// Package log provids glog functions.
package log

import (
	"path"
	"runtime"
	"strconv"

	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func reverse(a []any) []any {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
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

func Infof(fmt string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Infof(ctx, s+fmt, v...)
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

func Debugf(fmt string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Debugf(ctx, s+fmt, v...)
}

func Print(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log().Print(ctx, v...)
}

func Printf(fmt string, v ...any) {
	var ctx = context.TODO()
	// pc, file, line, _ := runtime.Caller(1)
	// name := runtime.FuncForPC(pc).Name()
	// s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	// v = append(v, s)
	// v = reverse(v)
	g.Log().Printf(ctx, fmt, v...)
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

func Warningf(fmt string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Warningf(ctx, s+fmt, v...)
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

func Errorf(fmt string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Errorf(ctx, s+fmt, v...)
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

func Panicf(fmt string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Panicf(ctx, s+fmt, v...)
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

func Fatalf(fmt string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Fatalf(ctx, s+fmt, v...)
}
