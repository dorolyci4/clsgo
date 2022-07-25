// Doc more on https://pkg.go.dev/github.com/spf13/viper
// https://yaml.org/ https://yaml.org/spec/1.2.2/

package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"time"
)

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

func ClsConfig(filename string /*Config file name*/, projectname string, monitoring bool) (cfg *viper.Viper, err error) {
	ViperInstance := viper.New()
	// viper could guess the extension of filename
	extension := path.Ext(filename)
	if extension != "" {
		ViperInstance.SetConfigType(extension[1:]) // REQUIRED if the config file does not have the extension in the name
	}
	var paths []string
	ViperInstance.SetConfigName(filename)                               // name of config file (without extension)
	clsDefConfigSearchPath(ViperInstance, paths, "/etc/"+projectname)   // path to look for the config file in
	clsDefConfigSearchPath(ViperInstance, paths, "$HOME/."+projectname) // call multiple times to add many search paths
	clsDefConfigSearchPath(ViperInstance, paths, ".")                   // optionally look for config in the working directory
	clsDefConfigSearchPath(ViperInstance, paths, "./config")            // optionally look for config in the working directory

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
