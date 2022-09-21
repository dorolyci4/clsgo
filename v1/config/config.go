// Package config provides functions while implemented by Viper and gcfg.
// See more on https://pkg.go.dev/github.com/spf13/viper
package config

import (
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Config = *viper.Viper

const defaultFilename = "config"

// WARNING:
// If the default configuration file does not exist when the program is running,
// even if it is created at runtime, the hot update function cannot be enabled.
// To use the hot update feature, you must recreate an instance and load it once.
// In test cases and CLSMT, configuration files are not used by default.
var Cfg Config

func init() {
	// Load default config
	Cfg = New(defaultFilename, "")
}

func matchFile(cfg Config, project, filename string) {
	cfg.SetConfigName(filename)
	// name of config file (without extension)
	// path to look for the config file in
	// call multiple times to add many search paths
	cfg.AddConfigPath(".")
	cfg.AddConfigPath("./config")
	if project != "" {
		cfg.AddConfigPath("/etc/" + project)
		cfg.AddConfigPath("$HOME/." + project)
	} else {
		cfg.AddConfigPath("/etc/")
		cfg.AddConfigPath("$HOME/")
	}
}

// If you want to use the default configuration file, you need to create it manually
func CreateDefault(project string) {
	instance := viper.New()
	matchFile(instance, project, defaultFilename)
	useDefaultValue(instance, project)

	instance.SetConfigType("yaml")
	instance.SafeWriteConfig()
}

func Default() Config {
	reload := viper.New()
	matchFile(reload, "", defaultFilename)
	// Find and read the config file
	reload.ReadInConfig()
	return reload
}

// Create a new configuration object.
// Filename is the name of the configuration file without a suffix,
// project is the name of the project corresponding to the configuration file,
// create determines whether to create the configuration file if it does not exist
func New(filename, project string) Config {
	instance := viper.New()
	matchFile(instance, project, filename)
	// Find and read the config file
	err := instance.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		// In some typical scenarios, CREATE should be false:
		// for terminal programs. Or, the configuration file is not required
		useDefaultValue(instance, project)
	}

	instance.OnConfigChange(func(in fsnotify.Event) {
		if in.Op&fsnotify.Write != 0 {
			instance.ReadInConfig()
		}
	})
	// Only works(enable watcher) when filename found out
	if instance.ReadInConfig() == nil {
		instance.WatchConfig()
	}
	return instance
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetIntWithDefault(cfg string, def int) int {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToInt(r.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetStringWithDefault(cfg string, def string) string {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToString(r.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetDurationWithDefault(cfg string, def time.Duration) time.Duration {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToDuration(r.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetBoolWithDefault(cfg string, def bool) bool {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToBool(r.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetFloat32WithDefault(cfg string, def float32) float32 {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToFloat32(r.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetFloat64WithDefault(cfg string, def float64) float64 {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToFloat64(r.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetIntSliceWithDefault(cfg string, def []int) []int {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToIntSlice(r.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetStringSliceWithDefault(cfg string, def []string) []string {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToStringSlice(r.Get(cfg))
}

// GetWithDefault return the match result form config file,
// Retrun default value(def) if not found
func GetInt64WithDefault(cfg string, def int64) int64 {
	r := Default()
	if !r.InConfig(cfg) {
		return def
	}
	return cast.ToInt64(r.Get(cfg))
}
