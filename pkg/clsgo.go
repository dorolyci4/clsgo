// Package log rely on package config,
// package config rely on package utils
package clsgo

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lovelacelee/clsgo/pkg/config"
)

var Cfg = config.Cfg

func init() {
}

// Goframe type map
type Map = g.Map
