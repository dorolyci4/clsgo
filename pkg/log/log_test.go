package log

import (
	"fmt"
)

// Exmaple
func ExampleClsLog() {
	var Log = ClsLog(&ClsFormatter{
		Prefix:      true,
		ForceColors: false,
	})
	fmt.Println(Log.Level)
}
