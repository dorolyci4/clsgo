package clsgo

import (
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
)

var Cfg config.Config

func init() {
	var err error
	// Load global config file
	Cfg, err = config.ClsConfig("config", "clsgo", true)
	if err != nil {
		log.Error("Config load failed!")
	}
	log.Info(Cfg.Get("project.name"))
}
