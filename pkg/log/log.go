// Package log provides glog functions.
package log

import (
	"context"
	"path"
	"runtime"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"

	// "github.com/gogf/gf/v2/util/gutil"
	// "github.com/gogf/gf/v2/errors/gcode"
	// "github.com/gogf/gf/v2/errors/gerror"
	"strings"

	"github.com/lovelacelee/clsgo/pkg/config"
)

var levelStringMap = map[string]int{
	"ALL":      glog.LEVEL_DEBU | glog.LEVEL_INFO | glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"DEV":      glog.LEVEL_DEBU | glog.LEVEL_INFO | glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"DEVELOP":  glog.LEVEL_DEBU | glog.LEVEL_INFO | glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"PROD":     glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"PRODUCT":  glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"DEBU":     glog.LEVEL_DEBU | glog.LEVEL_INFO | glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"DEBUG":    glog.LEVEL_DEBU | glog.LEVEL_INFO | glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"INFO":     glog.LEVEL_INFO | glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"NOTI":     glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"NOTICE":   glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"WARN":     glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"WARNING":  glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"ERRO":     glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"ERROR":    glog.LEVEL_ERRO | glog.LEVEL_CRIT,
	"CRIT":     glog.LEVEL_CRIT,
	"CRITICAL": glog.LEVEL_CRIT,
}

func init() {
	// glog functions could not follow my heart
	// Use runtime instead
	// g.Log().SetFlags(glog.F_TIME_STD | glog.F_CALLER_FN | glog.F_FILE_SHORT)

	// IMPORTANT: config load only after main called, not valid in go test
	// so here load them use viper manually
	if g.Log().GetConfig().Path == "" {
		Infoi("Need manual load log config")
		loadLogConfig("logger")
		loadLogConfig("logger.clsgo")
	}
}

func loadLogConfig(logger string) *glog.Config {
	var logLevel int
	if level, ok := levelStringMap[strings.ToUpper(config.Cfg.GetString(logger+".level"))]; ok {
		logLevel = level
	} else {
		logLevel = levelStringMap["ALL"]
	}

	// TODO: load to map, then convert to struct because of the map can use non-case-sensitive
	c := glog.Config{
		Path:         config.Cfg.GetString(logger + ".path"),
		File:         config.Cfg.GetString(logger + ".file"),
		Level:        logLevel,
		Prefix:       config.Cfg.GetString(logger + ".prefix"),
		StSkip:       config.Cfg.GetInt(logger + ".StSkip"),
		StStatus:     config.Cfg.GetInt(logger + ".StStatus"),
		StFilter:     config.Cfg.GetString(logger + ".StFilter"),
		RotateSize:   gfile.StrToSize(gconv.String(config.Cfg.GetString(logger + ".rotateSize"))),
		RotateExpire: gconv.Duration(config.Cfg.GetString(logger + ".rotateExpire")),
		// Db:              g.Cfg().GetInt(configpath + ".db"),
		// Pass:            g.Cfg().GetString(configpath + ".pass"),
		// MinIdle:         g.Cfg().GetInt(configpath + ".minIdle"),
		// MaxIdle:         g.Cfg().GetInt(configpath + ".maxIdle"),
		// MaxActive:       g.Cfg().GetInt(configpath + ".maxActive"),
		// IdleTimeout:     g.Cfg().GetDuration(configpath + ".idleTimeout"),
		// MaxConnLifetime: g.Cfg().GetDuration(configpath + ".maxConnLifetime"),
		// WaitTimeout:     g.Cfg().GetDuration(configpath + ".waitTimeout"),
		// DialTimeout:     g.Cfg().GetDuration(configpath + ".dialTimeout"),
		// ReadTimeout:     g.Cfg().GetDuration(configpath + ".readTimeout"),
		// WriteTimeout:    g.Cfg().GetDuration(configpath + ".writeTimeout"),
		// MasterName:      g.Cfg().GetString(configpath + ".masterName"), //Used in Sentinel mode
		// TLS:             g.Cfg().GetBool(configpath + ".tls"),
		// TLSSkipVerify:   g.Cfg().GetBool(configpath + ".tlsSkipVerify"),
	}

	Infoi(c)
	return &c
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
	s := "[" + path.Base(file) + ":" + strconv.Itoa(line) + " " + path.Base(name) + "]"

	g.Log().Fatalf(ctx, s+" "+f, v...)
}
