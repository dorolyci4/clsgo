// Package config provides functions while implemented by Viper and gcfg.
// See more on https://pkg.go.dev/github.com/spf13/viper
package config

import (
	"log"
	"os"
	// "path"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/lovelacelee/clsgo/pkg/utils"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Config = *viper.Viper

var Cfg Config

func init() {
	var generateDeafult = false
	// Load global config file
	Cfg = ClsConfig("config", "clsgo", generateDeafult)
}

// viper.ConfigWatch is not reliable
func monitor(cfg *viper.Viper) {
	for {
		time.Sleep(time.Second * 5) // delay
		err := cfg.ReadInConfig()   // Find and read the config file
		if err != nil {
			log.Printf("%v", err)
		}
	}
}

func clsDefConfigSearchPath(v *viper.Viper, paths []string, path string) []string {
	v.AddConfigPath(path) // optionally look for config in the working directory
	return append(paths, path)
}

// Param: <monitoring> will start a routine to watch file changes, and reload it. and the goroutine never ends.
// If <monitoring> is true and <filename> config file not exist, it will be create by default.
// <filename> does not include extension.
func ClsConfig(filename string /*Config file name without extension*/, projectname string, monitoring bool) (cfg *viper.Viper) {
	ViperInstance := viper.New()

	// open a goroutine to watch remote changes forever
	if monitoring {
		go monitor(ViperInstance)
	}
	// viper could guess the extension of filename
	// extension := path.Ext(filename)
	// if extension != "" {
	// 	ViperInstance.SetConfigType(extension[1:]) // REQUIRED if the config file does not have the extension in the name
	// }
	var paths = make([]string, 0)
	ViperInstance.SetConfigName(filename)                                       // name of config file (without extension)
	paths = clsDefConfigSearchPath(ViperInstance, paths, ".")                   // path to look for the config file in
	paths = clsDefConfigSearchPath(ViperInstance, paths, "./config")            // call multiple times to add many search paths
	paths = clsDefConfigSearchPath(ViperInstance, paths, "/etc/"+projectname)   // optionally look for config in the working directory
	paths = clsDefConfigSearchPath(ViperInstance, paths, "$HOME/."+projectname) // optionally look for config in the working directory
	// Test cases
	paths = clsDefConfigSearchPath(ViperInstance, paths, utils.GetCurrentAbPath()+"/../../config")

	err := ViperInstance.ReadInConfig() // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if monitoring {
				log.Printf("Config file not found in %v, default generated and used.\n", paths)
				ViperInstance.SetConfigType("yaml")
			}
			// Use default
			if errDef := ClsConfigDefault(ViperInstance, monitoring); errDef != nil {
				log.Println(errDef)
			}
			return ViperInstance
		} else {
			log.Fatalf("%v", err)
			os.Exit(1)
		}
	}
	return ViperInstance
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetIntWithDefault(cfg string, def int) int {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToInt(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetStringWithDefault(cfg string, def string) string {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToString(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetDurationWithDefault(cfg string, def time.Duration) time.Duration {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToDuration(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetBoolWithDefault(cfg string, def bool) bool {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToBool(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetFloat32WithDefault(cfg string, def float32) float32 {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToFloat32(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetFloat64WithDefault(cfg string, def float64) float64 {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToFloat64(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetIntSliceWithDefault(cfg string, def []int) []int {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToIntSlice(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetStringSliceWithDefault(cfg string, def []string) []string {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToStringSlice(Cfg.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetInt64WithDefault(cfg string, def int64) int64 {
	if !Cfg.InConfig(cfg) {
		return def
	}
	return cast.ToInt64(Cfg.Get(cfg))
}

// Functions implemented using goframe

// Get using gcfg
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
