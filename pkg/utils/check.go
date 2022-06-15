package utils

import (
	"fmt"
	"os"
	"strings"

	. "github.com/lovelacelee/clsgo/pkg/log"
)

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an app.
func ExitIfCheckArgsFailed(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func ExitIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
