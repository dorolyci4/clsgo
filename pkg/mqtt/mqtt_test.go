package mqtt_test

import (
	// "github.com/lovelacelee/clsgo/pkg/log"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/lovelacelee/clsgo/pkg/mqtt"
)

func Test(t *testing.T) {
	Example()
}

func Example() {
	mqtt := mqtt.New("tcp://192.168.137.100:1883", "lee", "lovelace", "ID_CLIENT_TEST")

	mqtt.Subscribe("test/#", 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			mqtt.Publish("test/abc", 2, true, "Hello"+strconv.Itoa(i))
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Wait()
}
