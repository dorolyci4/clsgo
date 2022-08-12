// Package log rely on package config,
// package config rely on package utils
package clsgo

import (
	"github.com/gogf/gf"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
)

var Cfg = config.Cfg

func init() {
	log.Warningfi("Version: %v", Version)
	log.Warningfi("Config Version: %v", Cfg.GetString("project.version"))
	log.Warningfi("GF-Version: %v", gf.VERSION)
}

// Goframe type map
type Map = g.Map
