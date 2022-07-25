package log_test

import (
	"github.com/lovelacelee/clsgo/pkg/log"
)

// Exmaple
func ExampleClsLog() {
	var Log = log.ClsLog(&log.Formatter)
	Log.Info(Log.Level)
}
