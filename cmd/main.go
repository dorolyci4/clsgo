/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-01-14 08:59:04
 * @LastEditTime    : 2022-06-28 19:18:38
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/main.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package main

import (
	// "sync"
	"time"

	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
)

func main() {

	// var workGroup sync.WaitGroup
	// workGroup.Add(1)

	cfg, err := config.Instance("config", "clsgo", true)
	if err != nil {
		log.Error("Config load failed!")
	}
	for {
		time.Sleep(time.Second * 1)
		log.Info("%v", cfg.Sub("log").GetString("prefix"))
	}

	// workGroup.Wait()
}
