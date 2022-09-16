package utils

import (
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

var PBTheme = map[string]progressbar.Option{
	"=>": progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}),
	"█": progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]█[reset]",
		SaucerHead:    "[green] [reset]",
		SaucerPadding: " ",
		BarStart:      "|",
		BarEnd:        "|",
	}),
}

// If max == -1, enable spinner progress bar,
// theme if one of PBTheme key
func ProgressBar(max int64, spin int, theme string, descPrefix ...string) *progressbar.ProgressBar {
	var prefix string
	for _, v := range descPrefix {
		prefix += v + " "
	}
	spin = func(n int) int {
		if n > 75 || n < 0 {
			return 0
		} else {
			return n
		}
	}(spin)
	return progressbar.NewOptions64(max,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		func(p string) progressbar.Option {
			if !IsEmpty(p) {
				return progressbar.OptionSetDescription(prefix)
			} else {
				return progressbar.OptionSpinnerType(spin)
			}
		}(prefix),
		func(p string) progressbar.Option {
			if !IsEmpty(PBTheme[theme]) {
				return PBTheme[theme]
			} else {
				return PBTheme["=>"]
			}
		}(theme),
	)
}
