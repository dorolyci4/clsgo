package log_test

import (
	"github.com/lovelacelee/clsgo/pkg/log"
)

// Exmaple
func Example() {
	log.Importi("Internal important info.")
	log.Errori("Internal error info.")
	log.Infoi("Internal info.")
	log.Warni("Internal warning info.")
	log.Print("Print message")
}
