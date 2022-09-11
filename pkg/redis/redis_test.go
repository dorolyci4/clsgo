package redis_test

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/pkg/redis"
	"github.com/lovelacelee/clsgo/pkg/version"
	"sync"
	"testing"
)

var workGroup sync.WaitGroup

const messageCount = 1000

func Test(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Run("New-Close", func(to *testing.T) {
			c := redis.New("default")
			defer c.Close()
			t.AssertNE(c, nil)
		})
	})

	// workGroup.Add(2)
	// go ExampleClient_Subscribe()
	// go ExampleClient_publish()
	// workGroup.Wait()
}

func ExampleNew() {
	c := redis.New("default")
	defer c.Close()
}

func ExampleClient_Do() {
	c := redis.New("default")
	defer c.Close()

	c.Do("SET", "clsgo", version.Version)
	rv, _ := c.Do("GET", "clsgo")
	fmt.Println(rv.String())
}

func ExampleClient_publish() {
	c := redis.New("default")
	defer c.Close()
	for i := 0; i < messageCount; i++ {
		c.Do("PUBLISH", "channel", "test")
	}
	workGroup.Done()
}

func ExampleClient_Subscribe() {
	c := redis.New("default")
	defer c.Close()
	notify := c.Subscribe("channel")
	var payload string
	for i := 0; i < messageCount; i++ {
		resp := <-notify
		channelRes := struct {
			Channel      string
			Pattern      string
			Payload      string
			PayloadSlice string
		}{}
		resp.Struct(&channelRes)
		payload = channelRes.Payload
	}
	fmt.Printf("Receive %v %v times\n", payload, messageCount)
	workGroup.Done()
	// Output
	// v0.0.9
	// Receive test 1000 times
}
