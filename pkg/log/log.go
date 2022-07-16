/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-01-14 09:01:57
 * @LastEditTime    : 2022-07-16 17:06:14
 * @LastEditors     : Lovelace
 * @Description     : Use logrus as the file logger
 * @FilePath        : /pkg/log/log.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 * // Init logger
	go func() {
		for {
			log.Update(cfg.Sub("log"))
			time.Sleep(time.Second)
		}
	}()
	ClsLog := log.ClsLog()

	for {
		time.Sleep(time.Second * 1)
		ClsLog.WithField("name", "lee").Trace("trace")
		ClsLog.WithField("name", "lee").Debug("debug")
		ClsLog.WithField("name", "lee").Info("info")
		ClsLog.WithField("name", "lee").Warn("warning")
		ClsLog.WithField("name", "lee").Error("error")
	}
*/
// log.go
package log

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// DIY logrus formatter
type ClsFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool
	// Override coloring based on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/
	EnvironmentOverrideColors bool
	// Force disabling colors.
	DisableColors bool
	Prefix        bool
}

// Unmarshal need TWO most important features:
// 1.struct member must be exportable
// 2.struct member name must be same as the description

type LogConf struct {
	Enable          bool   `yml:"enable"`
	SeparateFile    bool   `yml:"separateFile"`
	FilePath        string `yml:"filePath"`
	FileName        string `yml:"fileName"`
	LogLevel        string `yml:"logLevel"`
	RotateMode      string `yml:"rotateMode"`
	RotateMaxAge    uint   `yml:"rotateMaxage"`
	RotateTime      uint   `yml:"rotateTime"`
	RotateSaveCount uint   `yml:"rotateSaveCount"`
	RotateSize      uint   `yml:"rotateSize"`
}

// The global instance
var logrusInstance = logrus.New()

const RotateSizeMB = 1024 * 1024

// Global config of logrusInstance
var Config LogConf

func (f *ClsFormatter) isColored() bool {
	isColored := f.ForceColors

	if f.EnvironmentOverrideColors {
		switch force, ok := os.LookupEnv("CLICOLOR_FORCE"); {
		case ok && force != "0":
			isColored = true
		case ok && force == "0", os.Getenv("CLICOLOR") == "0":
			isColored = false
		}
	}

	return isColored && !f.DisableColors
}

func (f *ClsFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	var logString string
	_, color := loglevel(entry.Level.String())
	if f.Prefix {
		if f.isColored() && (entry.Logger.Out == os.Stderr || entry.Logger.Out == os.Stdout) {
			logString = fmt.Sprintf("%s[%s]%s[%s]", color, strings.ToUpper(entry.Level.String())[:4], ANSI_RESET, timestamp)
		} else {
			logString = fmt.Sprintf("[%s][%s]", strings.ToUpper(entry.Level.String())[:4], timestamp)
		}
		if entry.HasCaller() {
			fName := path.Base(entry.Caller.File)
			fFunc := path.Base(entry.Caller.Function)
			logString += fmt.Sprintf("[%s:%d %s]", fName, entry.Caller.Line, fFunc)
		}
		for k, v := range entry.Data {
			logString += fmt.Sprintf("[%s=%s]", k, v)
		}
	}
	logString += fmt.Sprintf("%s\n", entry.Message)
	b.WriteString(logString)
	return b.Bytes(), nil
}

func init() {
	logrusInstance.SetReportCaller(true)
	logrusInstance.SetFormatter(&ClsFormatter{
		Prefix:      false,
		ForceColors: true,
	})
	logrusInstance.SetLevel(logrus.TraceLevel)
}

// Unit of maxage is Hour
func writer(level string) *rotatelogs.RotateLogs {
	var logFullPath string
	var rotatelogs_suffix string
	var options []rotatelogs.Option
	logfile := path.Join(Config.FilePath, Config.FileName)
	if Config.RotateMode == "count" {
		rotatelogs_suffix = ".%Y%m%d"
		options = append(options, rotatelogs.WithMaxAge(-1))
		options = append(options, rotatelogs.WithRotationCount(Config.RotateSaveCount))
		options = append(options, rotatelogs.WithRotationSize((int64)(RotateSizeMB*Config.RotateSize)))
	} else {
		rotatelogs_suffix = "-%Y%m%d%H%M%S"
		options = append(options, rotatelogs.WithMaxAge(time.Hour*time.Duration(Config.RotateMaxAge)))
		options = append(options, rotatelogs.WithRotationTime(time.Hour*time.Duration(Config.RotateTime)))
	}
	if Config.SeparateFile {
		logFullPath = logfile + "-" + strings.ToUpper(level) + rotatelogs_suffix
	} else {
		logFullPath = logfile + rotatelogs_suffix
	}

	// Warning: logFullPath must have right date format, otherwise rotatelogs cannot remove expired files.
	wr, err := rotatelogs.New(
		logFullPath,
		options...,
	)

	if err != nil {
		panic(err)
	}
	return wr
}

func newhook() (hook logrus.Hook) {
	writeMap := lfshook.WriterMap{
		logrus.TraceLevel: writer("trace"),
		logrus.DebugLevel: writer("debug"),
		logrus.InfoLevel:  writer("info"),
		logrus.WarnLevel:  writer("warning"),
		logrus.ErrorLevel: writer("error"),
		logrus.FatalLevel: writer("fatal"),
		logrus.PanicLevel: writer("panic"),
	}

	lfHook := lfshook.NewHook(writeMap, &ClsFormatter{})

	return lfHook
}

func loglevel(loglevel string) (level logrus.Level, color string) {
	switch loglevel {
	case "trace":
		level = logrus.TraceLevel
		color = ANSI_DEFAULT
	case "debug":
		level = logrus.DebugLevel
		color = ANSI_WHITE
	case "info":
		level = logrus.InfoLevel
		color = ANSI_GREEN
	case "warning":
		level = logrus.WarnLevel
		color = ANSI_YELLOW
	case "error":
		level = logrus.ErrorLevel
		color = ANSI_RED
	case "fatal":
		level = logrus.FatalLevel
		color = ANSI_CYAN
	case "panic":
		level = logrus.PanicLevel
		color = ANSI_CYAN
	default:
		level = logrus.InfoLevel
		color = ANSI_DEFAULT
	}
	return
}

func Update(logcfg *viper.Viper) (logger *logrus.Logger, err error) {
	if logcfg == nil {
		err := errors.New("log config section is nil")
		Error("%v", err)
		return nil, err
	}
	logcfg.Unmarshal(&Config)

	if !Config.Enable {
		return logrusInstance, err
	}
	// Check if the path exist, create it first
	dir := path.Dir(path.Join(Config.FilePath, Config.FileName))
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			logrusInstance.Warnf("Create directory failed, err:%v\n", err)
		}
	}
	lvl, _ := loglevel(Config.LogLevel)
	logrusInstance.SetLevel(lvl)
	hooks := make(logrus.LevelHooks)

	hooks.Add(newhook())
	logrusInstance.ReplaceHooks(hooks)

	return logrusInstance, err
}

func ClsLog() *logrus.Logger {
	return logrusInstance
}
