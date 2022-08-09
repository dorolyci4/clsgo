package main

import (
	"sync"
)

func main() {
	var workGroup sync.WaitGroup
	workGroup.Add(1)
	App()
	workGroup.Wait()
}
