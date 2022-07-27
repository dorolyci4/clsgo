package config_test

import (
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
)

func ExampleClsConfig() {
	cfg, err := config.ClsConfig("config", "clsgo", true)
	if err != nil {
		log.Error("Config load failed!")
	}
	log.Info(cfg.Get("project.name"))
}
