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

const messageCount = 100

func Test(t *testing.T) {
	workGroup.Add(1)
	go ExampleClient_Push()
	workGroup.Add(1)
	go ExampleClient_Consume_cancel()
	workGroup.Wait()
}

func ExampleClient_Push() {
	exchange := mq.Exchange{
		ExchangeName: "clsgo-exchange",
		ExchangeType: "direct",
		Durable:      false,
		Internal:     false,
		AutoDelete:   true,
		Nowait:       false,
	}
	queue := mq.Queue{
		QueueName:  "clsgo-queue",
		Exclusive:  false,
		Durable:    false,
		AutoDelete: true,
		Nowait:     false,
	}
	addr := clsgo.Cfg.GetString("rabbitmq.server")
	queueClient := mq.New(addr, exchange, queue, "clsgo", "clsgo")
	if <-queueClient.Connected {
		defer queueClient.Close()
		// Make sure the client is ready
		for queueClient.IsReady == false {
			time.Sleep(time.Microsecond * 50)
		}
		message := []byte("message")
		// Attempt to push 10 message, one every 10 microseconds
		for i := 0; i < messageCount; i++ {
			time.Sleep(time.Microsecond * 10)
			if err := queueClient.Push(message); err != nil {
				fmt.Printf("Push failed: %s\n", err)
			}
		}
	}
	workGroup.Done()
}

func ExampleClient_Consume_cancel() {
	exchange := mq.Exchange{
		ExchangeName: "clsgo-exchange",
		ExchangeType: "direct",
		Durable:      false,
		Internal:     false,
		AutoDelete:   true,
		Nowait:       false,
	}
	queue := mq.Queue{
		QueueName:  "clsgo-queue",
		Exclusive:  false,
		Durable:    false,
		AutoDelete: true,
		Nowait:     false,
	}
	addr := clsgo.Cfg.GetString("rabbitmq.server")
	queueClient := mq.New(addr, exchange, queue, "clsgo", "clsgo")
	if <-queueClient.Connected {
		defer queueClient.Close()
		// Make sure the client is ready
		for queueClient.IsReady == false {
			time.Sleep(time.Microsecond * 50)
		}
		msgChan, err := queueClient.Consume(false)
		defer queueClient.CancelConsume()
		if err != nil {
			fmt.Printf("Consume failed : %s\n", err)
			return
		}
		count := 0
		for message := range msgChan {
			message.Ack(false)
			count++
			if count > (messageCount - 1) {
				break
			}
		}
		fmt.Printf("Consume %d messages\n", count)
	}
	workGroup.Done()
}
