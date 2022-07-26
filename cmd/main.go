package main

import (
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
	"sync"
)

func main() {

	var workGroup sync.WaitGroup
	workGroup.Add(1)

	cfg, err := config.ClsConfig("config", "clsgo", true)
	if err != nil {
		l.Error("Config load failed!")
	}
	log.ClsLog(cfg)
	App()
	workGroup.Wait()
}
