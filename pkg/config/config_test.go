package config_test

import (
	"fmt"
	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
)

func ExampleClsConfig() {
	cfg, err := config.ClsConfig("config", "clsgo", true)
	if err != nil {
		log.Error("Config load failed!")
	}
	fmt.Print(cfg.Get("project.name"))

	// Output:
	// clsgo
}

// Use global unique config instance
func Example() {
	// import "github.com/lovelacelee/clsgo/pkg" // which as clsgo package
	var Cfg = clsgo.Cfg
	fmt.Print(Cfg.Get("project.name"))
	// Output:
	// clsgo
}
