package log_test

import (
	"github.com/lovelacelee/clsgo/pkg/log"
	"testing"
)

// Exmaple
func Example() {
	log.Importi("Internal important info.")
	log.Errori("Internal error info.")
	log.Infoi("Internal info.")
	log.Warni("Internal warning info.")
	log.Print("Print message")
	log.Printf("%s\n", "Print message")
	log.Info(1, 2, 3)
	log.Warning(1, 2, 3, 4)
}

func Test(t *testing.T) {
	Example()
}
