package mq_test

import (
	"fmt"
	"github.com/lovelacelee/clsgo/pkg"
	mq "github.com/lovelacelee/clsgo/pkg/rabbitmq"
	"sync"
	"testing"
	"time"
)

var workGroup sync.WaitGroup

func Test(t *testing.T) {
	workGroup.Add(1)
	go ExampleClient_Push()
	workGroup.Add(1)
	go ExampleClient_Consume_cancel()
	workGroup.Wait()
}

func ExampleClient_Push() {
	queueName := "clsgo_queue"
	addr := clsgo.Cfg.GetString("rabbitmq.server")
	queueClient := mq.New(queueName, addr)
	if queueClient != nil {
		defer queueClient.Close()
		// Make sure the client is ready
		for queueClient.IsReady == false {
			time.Sleep(time.Microsecond * 50)
		}
		message := []byte("message")
		// Attempt to push 10 message, one every 10 microseconds
		for i := 0; i < 10; i++ {
			time.Sleep(time.Microsecond * 10)
			if err := queueClient.Push(message); err != nil {
				fmt.Printf("Push failed: %s\n", err)
			} else {
				fmt.Printf("Push %d succeeded!\n", i)
			}
		}
	}
	workGroup.Done()
}

func ExampleClient_Consume_cancel() {
	queueName := "clsgo_queue"
	addr := clsgo.Cfg.GetString("rabbitmq.server")
	queueClient := mq.New(queueName, addr)
	if queueClient != nil {
		defer queueClient.Close()
		// Make sure the client is ready
		for queueClient.IsReady == false {
			time.Sleep(time.Microsecond * 50)
		}
		msgChan, err := queueClient.Consume("clsgo_consumer")
		defer queueClient.CancelConsume("clsgo_consumer")
		if err != nil {
			fmt.Printf("Consume failed : %s\n", err)
			return
		}
		count := 0
		for message := range msgChan {
			fmt.Printf("Consume %d %s\n", count, string(message.Body))
			message.Ack(false)
			count++
			if count > 9 {
				break
			}
		}
	}
	workGroup.Done()
}
