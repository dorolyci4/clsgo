package redis_test

import (
	"fmt"
	// "github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/redis"
	"sync"
	"testing"
	// "time"
)

var workGroup sync.WaitGroup

const messageCount = 100

func Test(t *testing.T) {
	ExampleNew()
}

func ExampleNew() {
	c := redis.New("default")
	defer c.Close()

	v, err := c.SET("test", "test")
	fmt.Println(v.String())
	fmt.Println(err)
}
