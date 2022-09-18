// Package config provides functions while implemented by Viper and gcfg.
// See more on https://pkg.go.dev/github.com/spf13/viper
package config

import (
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/lovelacelee/clsgo/v1/utils"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Config = *viper.Viper

// WARNING: this will be always nil before Init() called.
var Cfg Config

func init() {
	Init("")
}

// Leave it to the user to initialize the configuration file
func Init(project string) {
	// Do not generate by default
	var generateDeafult = false
	// Load global config file, generate if not exist
	Cfg = ClsConfig("config", project, generateDeafult)
}

// viper.ConfigWatch is not reliable
func monitor(cfg *viper.Viper) {
	for {
		time.Sleep(time.Second * 5) // delay
		err := cfg.ReadInConfig()   // Find and read the config file
		if err != nil {
			utils.ErrorWithoutHeader("%v", err)
		}
	}
}

// Param: <monitoring> will start a routine to watch file changes, and reload it.
// and the goroutine never ends.
// If <monitoring> is true and <filename> config file not exist, it will be create by default.
// <filename> does not include extension.
func ClsConfig(filename string, projectname string, monitoring bool) (cfg *viper.Viper) {
	ViperInstance := viper.New()

	// open a goroutine to watch remote changes forever
	if monitoring {
		go monitor(ViperInstance)
	}

	ViperInstance.SetConfigName(filename)
	// name of config file (without extension)
	// path to look for the config file in
	// call multiple times to add many search paths
	ViperInstance.AddConfigPath(".")
	ViperInstance.AddConfigPath("./config")
	ViperInstance.AddConfigPath("/etc/" + projectname)
	ViperInstance.AddConfigPath("$HOME/." + projectname)

	// Find and read the config file
	err := ViperInstance.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if monitoring {
				ViperInstance.SetConfigType("yaml")
			}
			// Use default
			if errDef := ClsConfigDefault(ViperInstance, projectname, monitoring); errDef != nil {
				utils.WarnWithoutHeader("%v", errDef)
			}
			return ViperInstance
		} else {
			utils.ErrorWithoutHeader("%v", err)
			return nil
		}
	}
	return ViperInstance
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetIntWithDefault(cfg string, def int) int {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToInt(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetStringWithDefault(cfg string, def string) string {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToString(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetDurationWithDefault(cfg string, def time.Duration) time.Duration {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToDuration(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetBoolWithDefault(cfg string, def bool) bool {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToBool(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetFloat32WithDefault(cfg string, def float32) float32 {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToFloat32(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetFloat64WithDefault(cfg string, def float64) float64 {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToFloat64(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetIntSliceWithDefault(cfg string, def []int) []int {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToIntSlice(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetStringSliceWithDefault(cfg string, def []string) []string {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToStringSlice(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetInt64WithDefault(cfg string, def int64) int64 {
	if Cfg == nil || !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToInt64(Cfg.Get(cfg))
}

// Functions implemented using goframe, gvar returned
func Get(pattern string, def ...interface{}) (x *gvar.Var) {
	var ctx = gctx.New()
	result, err := gcfg.Instance().Get(ctx, pattern)
	if err != nil {
		if len(def) > 0 {
			return gvar.New(def[0])
		} else {
			return gvar.New(nil)
		}
	} else {
		return result
	}
}
