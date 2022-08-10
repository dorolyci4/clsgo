package rabbitmq_test

import (
	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/log"
	mq "github.com/lovelacelee/clsgo/pkg/rabbitmq"
	"github.com/lovelacelee/clsgo/pkg/utils"
	"sync"
	"testing"
)

var workGroup sync.WaitGroup

const messageCount = 10000
const retryTimes = 100

func Test(t *testing.T) {
	workGroup.Add(1)
	go ExampleClient_Publish()
	workGroup.Add(1)
	go ExampleClient_Consume_cancel()
	workGroup.Wait()
}

func ExampleClient_Publish() {
	exchange := mq.Exchange{
		ExchangeName: "clsgo-exchange",
		ExchangeType: "direct",
		Durable:      true,
		Internal:     false,
		AutoDelete:   false,
		Nowait:       false,
	}
	queue := mq.Queue{
		QueueName:  "clsgo-queue",
		Exclusive:  false,
		Durable:    true,
		AutoDelete: false,
		Nowait:     false,
	}
	addr := clsgo.Cfg.GetString("rabbitmq.server")
	queueClient := mq.New(addr, exchange, queue, "clsgo")
	defer queueClient.Close()
	message := []byte("message")
	log.Info("Start push")
	for i := 0; i < messageCount; i++ {
		// Publish blocks
		if err := queueClient.Publish(
			mq.PubStruct{
				ContentType:  "text/plain",
				Body:         message,
				DeliveryMode: 2,
			}, retryTimes); err != nil {
			log.Errorfi("Push failed: %s\n", err)
			break
		}
	}
	log.Info("Push routine done")
	workGroup.Done()
}

func ExampleClient_Consume_cancel() {
	exchange := mq.Exchange{
		ExchangeName: "clsgo-exchange",
		ExchangeType: "direct",
		Durable:      true,
		Internal:     false,
		AutoDelete:   false,
		Nowait:       false,
	}
	queue := mq.Queue{
		QueueName:  "clsgo-queue",
		Exclusive:  false,
		Durable:    true,
		AutoDelete: false,
		Nowait:     false,
	}
	addr := clsgo.Cfg.GetString("rabbitmq.server")
	queueClient := mq.New(addr, exchange, queue, "clsgo", "clsgo")
	defer queueClient.Close()

	count := 0
	for {
		select {
		case status := <-queueClient.NotifyStatus:
			if status == mq.MQCONN_READY {
				log.Info("Start consume")
				msgChan, err := queueClient.Consume(false)
			NEXT:
				if queueClient.Status != mq.MQCONN_READY {
					continue
				}
				message := <-msgChan
				log.Info(message.Body)
				err = message.Ack(false)
				utils.InfoIfError(err)
				count++
				if count > (messageCount - 1) {
					goto Exit
				}
				goto NEXT
			}

		default:
			continue
		}
	}

Exit:
	log.Info("Consumer routine done")
	workGroup.Done()
}
