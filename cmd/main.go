package main

import (
	"sync"
	"time"

	"github.com/lovelacelee/clsgo/pkg/log"
)

func main() {

	var workGroup sync.WaitGroup
	workGroup.Add(1)

	log.NewGLC(log.InitOption{
		Path:     "logs/",
		Prefix:   "backend",
		Interval: time.Duration(time.Second * 30),
		Reserve:  time.Duration(time.Hour * 12),
	})

	workGroup.Wait()
}
