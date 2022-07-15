/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-28 16:05:32
 * @LastEditTime    : 2022-07-15 11:20:45
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /test/config_test.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package clsgo_test

import (
	"testing"

	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
)

func TestConfig(t *testing.T) {
	cfg, err := config.ClsConfig("config", "clsgo", false)
	if err == nil {
		l.Info(cfg.Sub("log").GetString("enable"))
	} else {
		t.Errorf("Not passed")
	}
	err = cfg.WriteConfigAs("config.json")
	if err != nil {
		log.Error("%v", err)
	}
}
