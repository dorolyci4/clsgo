/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-01-14 09:01:57
 * @LastEditTime    : 2022-06-29 19:40:16
 * @LastEditors     : Lovelace
 * @Description     : Use logrus as the file logger
 * @FilePath        : /pkg/log/log.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
// log.go
package log

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// The global instance
var Instance = logrus.New()

const RotateSize = 1024 * 1024 * 5

// DIY logrus formatter
type ClsFormatter struct {
}

func (m *ClsFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var logString string
	if entry.HasCaller() {
		fName := path.Base(entry.Caller.File)
		logString = fmt.Sprintf("[%7s][%s][%s:%d %s] %s\n",
			strings.ToUpper(entry.Level.String()),
			timestamp, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		logString = fmt.Sprintf("[%7s][%s] %s\n",
			strings.ToUpper(entry.Level.String()),
			timestamp, entry.Message)
	}

	b.WriteString(logString)
	return b.Bytes(), nil
}

func init() {
	Instance.SetReportCaller(true)
	Instance.SetFormatter(&ClsFormatter{})
	Instance.SetLevel(logrus.TraceLevel)
}

func writer(logfile string, level string, maxage time.Duration, save uint) *rotatelogs.RotateLogs {

	// var rotatelogs_suffix = ".%Y%m%d"
	var rotatelogs_suffix = "-%Y%m%d%H%M%S"
	logFullPath := logfile + "-" + strings.ToUpper(level) + rotatelogs_suffix

	// Warning: logFullPath must have right date format, otherwise rotatelogs cannot remove expired files.
	wr, err := rotatelogs.New(
		logFullPath,
		// MaxAge and RotationCount cannot be both set
		// By count
		// rotatelogs.WithMaxAge(-1),
		// rotatelogs.WithRotationCount(save),
		// rotatelogs.WithRotationSize(RotateSize),
		// By time
		rotatelogs.WithMaxAge(time.Minute),
		rotatelogs.WithRotationTime(time.Second*5),
	)

	if err != nil {
		panic(err)
	}
	return wr
}

func rotate_settings(logfile string, maxage time.Duration, count uint) (err error) {
	writeMap := lfshook.WriterMap{
		logrus.TraceLevel: writer(logfile, "trace", maxage, count),
		logrus.DebugLevel: writer(logfile, "debug", maxage, count),
		logrus.InfoLevel:  writer(logfile, "info", maxage, count),
		logrus.WarnLevel:  writer(logfile, "warning", maxage, count),
		logrus.ErrorLevel: writer(logfile, "error", maxage, count),
		logrus.FatalLevel: writer(logfile, "fatal", maxage, count),
		logrus.PanicLevel: writer(logfile, "panic", maxage, count),
	}

	lfHook := lfshook.NewHook(writeMap, &ClsFormatter{})
	Instance.AddHook(lfHook)

	return
}

func loglevel(loglevel string) (level logrus.Level) {
	switch loglevel {
	case "trace":
		level = logrus.TraceLevel
	case "debug":
		level = logrus.DebugLevel
	case "info":
		level = logrus.InfoLevel
	case "warning":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	case "fatal":
		level = logrus.FatalLevel
	case "panic":
		level = logrus.PanicLevel
	default:
		level = logrus.InfoLevel
	}
	return
}

func Init(logfile, level string, maxAge time.Duration, count uint) (err error) {
	// Check if the path exist, create it first
	dir := path.Dir(logfile)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			Instance.Warnf("Create directory failed, err:%v\n", err)
		}
	}
	Info("%v", level)
	Instance.Debugf("%s %v %v", level, maxAge, count)
	Instance.SetLevel(loglevel(level))
	rotate_settings(logfile, maxAge, count)

	return
}
