package log_test

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
)

// Example
func Example() {
	log.Debugi("Internal important info.")
	log.Errori("Internal error info.")
	log.Infoi("Internal info.")
	log.Warningi("Internal warning info.")
	log.Print("Print message")
	log.Printf("%s\n", "Print message")
	log.Info(1, 2, 3)
	log.Warning(1, 2, 3, 4)
	log.Red("red")
	log.Blue("blue")
	fmt.Println(log.BlueString("test"))
}

func Test(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Run("internal", func(_ *testing.T) {
			log.Print("Print")
			log.Printf("%s", "Printf")
			log.Ignore()
			log.Ignoref("")
			log.Infoi()
			log.Infofi("test %v %v %v", "1", 2, 3)
			log.Debugi("1", "2", "3", "4")
			log.Debugfi("")
			log.Warningi()
			log.Warningfi("")
			log.Errori()
			log.Errorfi("")
		})
		t.Run("logger", func(_ *testing.T) {
			log.Info()
			log.Infof("test %v %v %v", "1", 2, 3)
			log.Debug()
			log.Debugf("")
			log.Warning()
			log.Warningf("")
			log.Error()
			log.Errorf("")

		})
	})
}
