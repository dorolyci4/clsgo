// Package database wraps gorm for convenience, we don't use goframe-gdb because of
// it's commitment of NEVER-SUPPORT-MIGRATE-SPECIALITY.
package database

// Package db provides gorm wrappers to clsgo
import (
	"github.com/lovelacelee/clsgo/v1/config"
)

type DbConfig struct {
	Dsn   string
	Type  string
	Redis string
}

var Config map[string]DbConfig

func loadConfig() {
	config.Cfg.UnmarshalKey("database", &Config)
	// viper := clsgo.Cfg.Sub("database")
	// viper.Unmarshal(&Config)
}

func GetConfig(group string) DbConfig {
	return Config[group]
}

func init() {
	loadConfig()
}
