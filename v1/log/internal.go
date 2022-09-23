// Internal log provides *i internal functions.
// Only output to stdout

package log

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"path"
	"runtime"
	"strconv"
)

func Print(v ...any) {
	fmt.Println(v...)
}

func Printf(f string, v ...any) {
	fmt.Printf(f, v...)
}

func Ignore(v ...any) {
	return
}

func Ignoref(fmt string, v ...any) {
	return
}

func Infoi(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log("clsgo").Info(ctx, v...)
}

func Infofi(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log("clsgo").Infof(ctx, s+" "+f, v...)
}

func Debugi(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log("clsgo").Debug(ctx, v...)
}

func Debugfi(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log("clsgo").Debugf(ctx, s+" "+f, v...)
}
func Warningi(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log("clsgo").Warning(ctx, v...)
}

func Warningfi(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log("clsgo").Warningf(ctx, s+" "+f, v...)
}

func Errori(v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	v = append(v, s)
	v = reverse(v)
	g.Log("clsgo").Error(ctx, v...)
}

func Errorfi(f string, v ...any) {
	var ctx = context.TODO()
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log("clsgo").Errorf(ctx, s+" "+f, v...)
}

// func Panici(v ...any) {
// 	var ctx = context.TODO()
// 	pc, file, line, _ := runtime.Caller(1)
// 	name := runtime.FuncForPC(pc).Name()
// 	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
// 	v = append(v, s)
// 	v = reverse(v)
// 	g.Log("clsgo").Panic(ctx, v...)
// }

// func Panicfi(f string, v ...any) {
// 	var ctx = context.TODO()
// 	pc, file, line, _ := runtime.Caller(1)
// 	name := runtime.FuncForPC(pc).Name()
// 	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

// 	g.Log("clsgo").Panicf(ctx, s+" "+f, v...)
// }

// func Fatali(v ...any) {
// 	var ctx = context.TODO()
// 	pc, file, line, _ := runtime.Caller(1)
// 	name := runtime.FuncForPC(pc).Name()
// 	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
// 	v = append(v, s)
// 	v = reverse(v)
// 	g.Log("clsgo").Fatal(ctx, v...)
// }

// func Fatalfi(f string, v ...any) {
// 	var ctx = context.TODO()
// 	pc, file, line, _ := runtime.Caller(1)
// 	name := runtime.FuncForPC(pc).Name()
// 	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

// 	g.Log("clsgo").Fatalf(ctx, s+" "+f, v...)
// }
