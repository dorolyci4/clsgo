package internal

import (
	"github.com/spf13/viper"
)

func LoadWithDefault(c *viper.Viper, cfg string, def any) any {
	r := c.Get(cfg)
	if r != nil {
		return r
	}
	return def
}
