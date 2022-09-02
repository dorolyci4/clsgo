package config

func ClsConfigDefault(cfg Config, gen bool) error {
	cfg.SetDefault("project.name", "clsgo")
	cfg.SetDefault("project.seemore", "https://pkg.go.dev/github.com/lovelacelee/clsgo")

	cfg.SetDefault("logger", map[string]any{
		"path":                "logs/",
		"file":                "{Y-m-d}.log",
		"stStatus":            0,
		"level":               "all",
		"rotateSize":          "100MB",
		"rotateBackupLimit":   5,
		"rotateCheckInterval": "1m",
		"RotateBackupExpire":  "1d",
		"writerColorEnable":   true,
		"stdoutColorDisabled": false,
		"clsgo": map[string]any{
			"stStatus":            0,
			"level":               "all",
			"prefix":              "[CLSGO]",
			"writerColorEnable":   true,
			"stdoutColorDisabled": false,
		},
	})

	if gen {
		return cfg.SafeWriteConfig()
	} else {
		return nil
	}
}
