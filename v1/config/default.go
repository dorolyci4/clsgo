package config

import (
	"github.com/lovelacelee/clsgo/v1/utils"
)

// With all necessary nodes
func useDefaultValue(cfg Config, name string) {
	cfg.SetDefault("project.name", name)
	cfg.SetDefault("project.seemore", "https://pkg.go.dev/github.com/lovelacelee/clsgo")

	cfg.SetDefault("logger", map[string]any{
		"path":                utils.TempPath("logs/"),
		"file":                "{Y-m-d}.log",
		"stStatus":            0,
		"prefix":              "[" + name + "]",
		"level":               "all",
		"rotateSize":          "1MB",
		"rotateBackupLimit":   1,
		"rotateCheckInterval": "1m",
		"RotateBackupExpire":  "1d",
		"writerColorEnable":   true,
		"stdoutColorDisabled": false,
		"clsgo": map[string]any{
			"stStatus":            0,
			"prefix":              "[" + name + "]",
			"level":               "info",
			"writerColorEnable":   true,
			"stdoutColorDisabled": false,
		},
	})
}
