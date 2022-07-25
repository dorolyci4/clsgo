package config_test

import (
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
	"time"
)

func ExampleClsConfig() {
	cfg, err := config.ClsConfig("config", "clsgo", true)
	if err != nil {
		log.ClsLog(&log.Formatter).Error("Config load failed!")
	}
	// Log monitor
	go func() {
		for {
			log.Update(cfg.Sub("log"))
			time.Sleep(time.Second)
		}
	}()
}
