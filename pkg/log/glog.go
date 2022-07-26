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
	def.Path = config.GCfgGet("logger.path", "logs/").String()
	def.StdoutPrint = config.GCfgGet("logger.stdout", true).Bool()
	def.File = config.GCfgGet("logger.file", "GFrame-{Y-m-d}.log").String()
	def.Level = glog.LEVEL_ALL
	def.RotateSize = config.GCfgGet("logger.rotateSize", 5*1024*1024).Int64()
	def.RotateBackupLimit = config.GCfgGet("logger.rotateBackupLimit", 1).Int()
	//Default glog configuration
	glog.SetDefaultHandler(func(ctx context.Context, in *glog.HandlerInput) {
		in.Logger.SetConfig(def)
		in.Next(ctx)
	})
}

func Logger() interface{} {
	return g.Log()
}
