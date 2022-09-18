package mqtt_test

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/lovelacelee/clsgo/v1/config"
	"github.com/lovelacelee/clsgo/v1/utils"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/mqtt"
)

const PublishCount = 5

func clean() {
	utils.DeletePath("logs")
	utils.DeleteFiles(utils.Cwd(), "/*.yaml$")
	utils.DeleteFiles(utils.Cwd(), "/*.xml$")
}
func Test(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {

	})
	clean()
}

func Example() {
	server := config.GetStringWithDefault("mqtt.server", "tcp://192.168.137.100:1883")
	user := config.GetStringWithDefault("mqtt.user", "lee")
	password := config.GetStringWithDefault("mqtt.password", "lovelace")

	mqtt := mqtt.New(server, user, password, "ID_CLIENT_TEST")
	defer mqtt.Close()
	mqtt.Subscribe("test/#", 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		log.Info("Publish routine begin")
		for i := 0; i < PublishCount; i++ {
			mqtt.Publish("test/abc", 2, true, "Hello"+strconv.Itoa(i))
			time.Sleep(time.Microsecond)
		}
		wg.Done()
		log.Info("Publish routine done")
	}()
	go func() {
		log.Info("Sub routine begin")
		// delivery := mqtt.Consume()
		// for m := range delivery {
		// 	log.Info(m.Payload())
		// }
		wg.Done()
		log.Info("Sub routine done")
	}()
	wg.Wait()
}
