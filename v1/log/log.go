// Package log provides glog functions.
package log

import (
	"context"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/lovelacelee/clsgo/v1/config"
)

func init() {
	// glog functions could not follow my heart
	// Use runtime instead
	// g.Log().SetFlags(glog.F_TIME_STD | glog.F_CALLER_FN | glog.F_FILE_SHORT)

	// IMPORTANT: g.Cfg() config load only after main called, not valid in go test.
	// so here load them use viper manually
	if g.Log("clsgo").GetConfig().Path == "" {
		// Will generate logs path by default, sometimes you don't want this happen
		// such as in test cases, or used in command tool like clsmt.
		// g.Log().SetConfigWithMap(loadLogConfig("logger"))

		// Internal logger instance
		g.Log("clsgo").SetConfigWithMap(loadLogConfig("logger.clsgo"))
	}
}

func loadLogConfig(logger string) map[string]any {
	m := map[string]any{
		"flags":                glog.F_TIME_STD,
		"path":                 config.GetStringWithDefault(logger+".path", "logs/"),
		"file":                 config.GetStringWithDefault(logger+".file", "{Y-m-d}.log"),
		"level":                config.GetStringWithDefault(logger+".level", "dev"),
		"prefix":               config.GetStringWithDefault(logger+".prefix", ""),
		"stSkip":               config.GetIntWithDefault(logger+".StSkip", 0),
		"stStatus":             config.GetIntWithDefault(logger+".StStatus", 0),
		"stFilter":             config.GetStringWithDefault(logger+".StFilter", ""),
		"header":               config.GetBoolWithDefault(logger+".header", false),
		"stdout":               config.GetBoolWithDefault(logger+".stdout", true),
		"rotateSize":           config.GetStringWithDefault(logger+".rotateSize", "1MB"),
		"rotateExpire":         config.GetStringWithDefault(logger+".rotateExpire", "0"),
		"rotateBackupLimit":    config.GetStringWithDefault(logger+".rotateBackupLimit", "1"),
		"rotateBackupExpire":   config.GetStringWithDefault(logger+".rotateBackupExpire", "0"),
		"rotateBackupCompress": config.GetStringWithDefault(logger+".rotateBackupCompress", "0"),
		"rotateCheckInterval":  config.GetStringWithDefault(logger+".rotateCheckInterval", "1m"),
		"stdoutColorDisabled":  config.GetBoolWithDefault(logger+".stdoutColorDisabled", false),
		"writerColorEnable":    config.GetBoolWithDefault(logger+".writerColorEnable", true),
	}
	return m
}

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
	if strings.HasPrefix(f, "[CLSGO[") {
		g.Log().Warningf(ctx, f[6:], v...)
		return
	}
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	g.Log().Infof(ctx, s+" "+f, v...)
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
	if strings.HasPrefix(f, "[CLSGO[") {
		g.Log().Warningf(ctx, f[6:], v...)
		return
	}
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	g.Log().Debugf(ctx, s+" "+f, v...)
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
	if strings.HasPrefix(f, "[CLSGO[") {
		g.Log().Warningf(ctx, f[6:], v...)
		return
	}
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	g.Log().Warningf(ctx, s+" "+f, v...)
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
	if strings.HasPrefix(f, "[CLSGO[") {
		g.Log().Warningf(ctx, f[6:], v...)
		return
	}
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	g.Log().Errorf(ctx, s+" "+f, v...)
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
	if strings.HasPrefix(f, "[CLSGO[") {
		g.Log().Warningf(ctx, f[6:], v...)
		return
	}
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	g.Log().Panicf(ctx, s+" "+f, v...)
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
	if strings.HasPrefix(f, "[CLSGO[") {
		g.Log().Warningf(ctx, f[6:], v...)
		return
	}
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"
	g.Log().Fatalf(ctx, s+" "+f, v...)
}
