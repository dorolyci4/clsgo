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
