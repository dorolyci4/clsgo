package db

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/lovelacelee/clsgo/pkg"
)

var Config gdb.Config

func loadConfig() gdb.Config {
	viper := clsgo.Cfg.Sub("database")
	viper.Unmarshal(&Config)
	return Config
}

func init() {
	gdb.SetConfig(loadConfig())
}
