package redis_test

import (
	"fmt"
	"github.com/lovelacelee/clsgo/pkg/redis"
	"sync"
	"testing"
)

var workGroup sync.WaitGroup

const messageCount = 1000

func Test(t *testing.T) {
	ExampleNew()
	ExampleClient_Do()
	workGroup.Add(2)
	go ExampleClient_Subscribe()
	go ExampleClient_publish()
	workGroup.Wait()
}

func ExampleNew() {
	c := redis.New("default")
	defer c.Close()
}

func ExampleClient_Do() {
	c := redis.New("default")
	defer c.Close()

	c.Do("SET", "test", "test")
	rv, _ := c.Do("GET", "test")
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
	fmt.Println("wait chan")
	for i := 0; i < messageCount; i++ {
		resp := <-notify
		channelRes := struct {
			Channel      string
			Pattern      string
			Payload      string
			PayloadSlice string
		}{}
		resp.Struct(&channelRes)
		fmt.Println(channelRes.Payload)
	}
	workGroup.Done()
}
