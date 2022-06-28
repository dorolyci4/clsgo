// Doc more on https://pkg.go.dev/github.com/spf13/viper
// https://yaml.org/ https://yaml.org/spec/1.2.2/

package config

import (
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

type Config = *viper.Viper

// viper.ConfigWatch is not reliable
func monitor(cfg Config) {
	for {
		time.Sleep(time.Second * 5) // delay
		err := cfg.ReadInConfig()   // Find and read the config file
		if err != nil {
			log.Error("%v", err)
		}
	}
}

func Instance(filename string, projectname string, monitoring bool) (cfg Config, err error) {
	ViperInstance := viper.New()
	// viper could guess the extension of filename
	extension := path.Ext(filename)
	if extension != "" {
		ViperInstance.SetConfigType(extension[1:]) // REQUIRED if the config file does not have the extension in the name
	}
	ViperInstance.SetConfigName(filename)                // name of config file (without extension)
	ViperInstance.AddConfigPath("/etc/" + projectname)   // path to look for the config file in
	ViperInstance.AddConfigPath("$HOME/." + projectname) // call multiple times to add many search paths
	ViperInstance.AddConfigPath(".")                     // optionally look for config in the working directory
	err = ViperInstance.ReadInConfig()                   // Find and read the config file
	if err != nil {                                      // Handle errors reading the config file
		log.Error("Config file(%v) load failed: %v", filename, err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Important("Make sure config file(%v) exist first.", filename)
		}
		os.Exit(1)
	}
	// open a goroutine to watch remote changes forever
	if monitoring {
		go monitor(ViperInstance)
	}
	return ViperInstance, err
}
