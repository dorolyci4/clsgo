package log

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/lovelacelee/clsgo/pkg/config"
)

// Warning: Clslog not available any more when Goframe log enabled
func UseGlog() {
	var def glog.Config
	def.Path = config.Get("logger.path", "logs/").String()
	def.StdoutPrint = config.Get("logger.stdout", true).Bool()
	def.File = config.Get("logger.file", "GFrame-{Y-m-d}.log").String()
	def.Level = glog.LEVEL_ALL
	def.RotateSize = config.Get("logger.rotateSize", 5*1024*1024).Int64()
	def.RotateBackupLimit = config.Get("logger.rotateBackupLimit", 1).Int()
	//Default glog configuration
	glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
		in.Logger.SetConfig(def)
		in.Next(ctx)
	})
}

func Logger() interface{} {
	return g.Log()
}
