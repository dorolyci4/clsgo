/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-13 14:46:44
 * @LastEditTime    : 2022-06-30 17:09:36
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/config/config.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
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

func ClsConfig(filename string, projectname string, monitoring bool) (cfg *viper.Viper, err error) {
	ViperInstance := viper.New()
	// viper could guess the extension of filename
	extension := path.Ext(filename)
	if extension != "" {
		ViperInstance.SetConfigType(extension[1:]) // REQUIRED if the config file does not have the extension in the name
	}
	var paths []string
	ViperInstance.SetConfigName(filename)              // name of config file (without extension)
	ViperInstance.AddConfigPath("/etc/" + projectname) // path to look for the config file in
	paths = append(paths, "/etc/"+projectname)
	ViperInstance.AddConfigPath("$HOME/." + projectname) // call multiple times to add many search paths
	paths = append(paths, "$HOME/."+projectname)
	ViperInstance.AddConfigPath(".") // optionally look for config in the working directory
	paths = append(paths, ".")
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
