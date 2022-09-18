package rabbitmq_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/config"
	"github.com/lovelacelee/clsgo/v1/log"
	mq "github.com/lovelacelee/clsgo/v1/rabbitmq"
	"github.com/lovelacelee/clsgo/v1/utils"
)

var workGroup sync.WaitGroup

const messageCount = 10
const retryTimes = 3

func clean() {
	utils.DeleteThingsInDir("logs")
	utils.DeletePath("logs")
	utils.DeleteFiles(utils.Cwd(), "/*.yaml$")
	utils.DeleteFiles(utils.Cwd(), "/*.xml$")
}
func Test(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Run("TestRabbitMQ", func(to *testing.T) {

		})
	})
	clean()
	// workGroup.Add(1)
	// go ExampleClient_Publish()
	// workGroup.Wait()
	// workGroup.Add(1)
	// go ExampleClient_Consume_cancel()
	// workGroup.Wait()
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
	addr := config.Cfg.GetString("rabbitmq.server")
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

	clean()
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
	addr := config.Cfg.GetString("rabbitmq.server")
	queueClient := mq.New(addr, exchange, queue, "clsgo", "clsgo")
	defer queueClient.Close()

	count := 0
	lastMessage := ""
	timeoutTimes := 0
	for {
		select {
		// If MQ connection or channel closed, client will reconnect automatically,
		// Here we just wait it be ready for consume
		case status := <-queueClient.NotifyStatus:
			if status == mq.MQCONN_READY {
				log.Info("Start consume")
				msgChan, err := queueClient.Consume(false)
				utils.InfoIfError(err, log.Errorf)
			NEXT: //Continous consume
				if queueClient.Status != mq.MQCONN_READY {
					continue
				}
				// message := <-msgChan
				message, err := utils.ReadChanWithTimeout(context.Background(), msgChan, 2*time.Second)
				if utils.InfoIfError(err, log.Errorf) != nil {
					timeoutTimes++
					if timeoutTimes > 5 {
						goto Exit
					}
				}
				lastMessage = string(message.Body)
				message.Ack(false)
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
	log.Infof("Consumer routine done: %d %s", count, lastMessage)
	workGroup.Done()

	clean()
}
