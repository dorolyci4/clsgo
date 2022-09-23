// Package database wraps gorm for convenience,
// clsgo don't use goframe-gdb because of
// it's commitment of NEVER-SUPPORT-MIGRATE-SPECIALITY.
package database

// Package db provides gorm wrappers to clsgo
import (
	"github.com/lovelacelee/clsgo/v1/config"
	"github.com/lovelacelee/clsgo/v1/utils"
)

type DbConfig struct {
	// mysql: "$user:$password@tcp($host:3306)/$dbname?charset=utf8&parseTime=True&loc=Local"
	// sqlite: "file::memory:?cache=shared" or just give the database filename "clsgo.db"
	Dsn string
	// "mysql","sqlite"
	Type string
}

var Config map[string]DbConfig

func loadConfig(cfg config.Config) {
	if cfg.InConfig("database") {
		cfg.UnmarshalKey("database", &Config)
	} else {
		Config = make(map[string]DbConfig)

		Config["default"] = DbConfig{
			Dsn:  "clsgo.db",
			Type: "sqlite",
		}
	}
}

// Default() make sure reload if any change to the config.yaml
func GetConfig(group ...string) DbConfig {
	loadConfig(config.Default())
	g := utils.Param(group, "default")
	return Config[g]
}

func init() {
	loadConfig(config.Cfg)
}
