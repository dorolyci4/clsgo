// Package Db wraps gorm for convenience, we don't use goframe-gdb because of
// it's commitment of NEVER-SUPPORT-MIGRATE-SPECIALITY
package db

import (
	"github.com/lovelacelee/clsgo/pkg"
)

type DbConfig struct {
	Dsn  string
	Type string
}

var Config map[string]DbConfig

func loadConfig() {
	clsgo.Cfg.UnmarshalKey("database", &Config)
	// viper := clsgo.Cfg.Sub("database")
	// viper.Unmarshal(&Config)
}

func GetConfig(group string) DbConfig {
	return Config[group]
}

func init() {
	loadConfig()
}
