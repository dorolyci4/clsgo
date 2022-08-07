package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Set up channel on which to send signal notifications.
// We must use a buffered channel or risk missing the signal
// if we're not ready to receive when the signal is sent.
func SetupExitNotify(done chan bool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		// Block until a signal is received.
		r := <-c
		done <- true
		fmt.Printf("Ctrl+C pressed in Terminal: %v\n", r.String())
	}()
}
