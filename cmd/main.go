/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-01-14 08:59:04
 * @LastEditTime    : 2022-06-29 18:37:10
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
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
	"path"
	"time"
)

type LogConf struct {
	Enable   bool   `yml:"enable"`
	FilePath string `yml:"filepath"`
	FileName string `yml:"filename"`
	LogLevel string `yml:"loglevel"`
	MaxAge   uint   `yml:"maxage"`
	Count    uint   `yml:"count"`
}

func main() {

	// var workGroup sync.WaitGroup
	// workGroup.Add(1)

	cfg, err := config.Instance("config", "clsgo", true)
	if err != nil {
		log.Error("Config load failed!")
	}
	// Init logger
	var config LogConf
	err = cfg.Sub("log").Unmarshal(&config)
	utils.WarnIfError(err)
	err = log.Init(
		path.Join(config.FilePath, config.FileName),
		config.LogLevel,
		//time.Duration(config.MaxAge)*time.Hour*24,
		time.Duration(config.MaxAge)*time.Second,
		config.Count)
	utils.WarnIfError(err)

	for {
		time.Sleep(time.Second * 1)
		log.Info("%v", cfg.Sub("log").GetString("level"))
		log.Instance.Trace("hello")
		log.Instance.WithField("name", "lee").Trace("test")
		log.Instance.WithField("name", "lee").Info("test")
		log.Instance.WithField("name", "lee").Warn("test")
		log.Instance.WithField("name", "lee").Debug("test")
		log.Instance.WithField("name", "lee").Error("test")
	}

	// workGroup.Wait()
}
