/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-28 16:05:32
 * @LastEditTime    : 2022-06-30 16:57:59
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /test/config_test.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package clsgo

import (
	"testing"

	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
)

func TestConfig(t *testing.T) {
	cfg, err := config.ClsConfig("config", "clsgo", false)
	if err == nil {
		log.Info("%v", cfg.Sub("log").GetString("enable"))
	} else {
		t.Errorf("Not passed")
	}
	err = cfg.WriteConfigAs("config.json")
	if err != nil {
		log.Error("%v", err)
	}
}
