// Package config provides functions while implemented by Viper and gcfg.
// See more on https://pkg.go.dev/github.com/spf13/viper
package config

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/lovelacelee/clsgo/pkg/utils"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"time"
)

type Config = *viper.Viper

var Cfg Config

func init() {
	var err error
	// Load global config file
	Cfg, err = ClsConfig("config", "clsgo", true)
	if err != nil {
		log.Println("Config load failed!")
	}
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

// Param: monitoring will use go routine to watch file changes, and reload it.
// the goroutine never ends.
func ClsConfig(filename string /*Config file name*/, projectname string, monitoring bool) (cfg *viper.Viper, err error) {
	ViperInstance := viper.New()
	// viper could guess the extension of filename
	extension := path.Ext(filename)
	if extension != "" {
		ViperInstance.SetConfigType(extension[1:]) // REQUIRED if the config file does not have the extension in the name
	}
	var paths = make([]string, 0)
	ViperInstance.SetConfigName(filename)                                       // name of config file (without extension)
	paths = clsDefConfigSearchPath(ViperInstance, paths, "/etc/"+projectname)   // path to look for the config file in
	paths = clsDefConfigSearchPath(ViperInstance, paths, "$HOME/."+projectname) // call multiple times to add many search paths
	paths = clsDefConfigSearchPath(ViperInstance, paths, ".")                   // optionally look for config in the working directory
	paths = clsDefConfigSearchPath(ViperInstance, paths, "./config")            // optionally look for config in the working directory
	paths = clsDefConfigSearchPath(ViperInstance, paths, utils.GetCurrentAbPath()+"/../../config")

	err = ViperInstance.ReadInConfig() // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Make sure config file(%v) exist in %v", filename, paths)
		}
		log.Fatalf("%v", err)
		os.Exit(1)
	}
	// open a goroutine to watch remote changes forever
	if monitoring {
		go monitor(ViperInstance)
	}
	return ViperInstance, err
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
