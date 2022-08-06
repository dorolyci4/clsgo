// Package mq exports a RabbitMQ Client object that wraps the official go library. It
// automatically reconnects when the connection fails, and
// blocks all pushes until the connection succeeds. It also
// confirms every outgoing message, so none are lost.
// It doesn't automatically ack each message, but leaves that
// to the parent process, since it is usage-dependent.
package mq

import (
	"context"
	"errors"
	"strings"
	"time"

	clsgo "github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Rabbit MQ client
type Client struct {
	host            string
	queueName       string
	connection      *amqp.Connection
	channel         *amqp.Channel
	done            chan bool
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	// When IsReady is true, everything is ok for push and consume
	IsReady bool
}

// Default retry params
var (
	// When reconnecting to the server after connection failure
	reconnectDelay = 5 * time.Second

	// When setting up the channel after a channel exception
	reInitDelay = 2 * time.Second

	// When resending messages the server didn't confirm
	resendDelay = 5 * time.Second
)

var (
	errNotConnected  = errors.New("not connected to a server")
	errAlreadyClosed = errors.New("already closed: not connected to the server")
	errShutdown      = errors.New("client is shutting down")
)

func init() {
	reconnect := clsgo.Cfg.GetInt("rabbitmq.reconnect")
	reinit := clsgo.Cfg.GetInt("rabbitmq.reinit")
	resend := clsgo.Cfg.GetInt("rabbitmq.resend")
	if reconnect > 0 {
		reconnectDelay = time.Duration(reconnect) * time.Second
	}
	if reinit > 0 {
		reInitDelay = time.Duration(reinit) * time.Second
	}
	if resend > 0 {
		resendDelay = time.Duration(resend) * time.Second
	}
	log.Infof("MQ reconnect:%v reinit:%v resend:%v", reconnect, reinit, resend)
}

// New creates a new consumer state instance, and automatically
// attempts to connect to the server.
func New(queueName, addr string) *Client {
	if !strings.Contains(addr, "@") {
		log.Error("MQ address invalid, use right formmat: amqp://test:test@localhost:5672/")
		return nil
	}
	client := Client{
		host:      strings.SplitAfter(addr, "@")[1],
		queueName: queueName,
		done:      make(chan bool),
	}
	go client.handleReconnect(addr)
	return &client
}

// handleReconnect will wait for a connection error on
// notifyConnClose, and then continuously attempt to reconnect.
func (client *Client) handleReconnect(addr string) {
	for {
		client.IsReady = false
		log.Info("Attempting to connect: ", client.host)

		conn, err := client.connect(addr)

		if err != nil {
			log.Error("Failed to connect. Retrying...")

			select {
			case <-client.done:
				return
			case <-time.After(reconnectDelay):
			}
			continue
		}

		if done := client.handleReInit(conn); done {
			break
		}
	}
}

// connect will create a new AMQP connection
func (client *Client) connect(addr string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(addr)

	if err != nil {
		return nil, err
	}

	client.changeConnection(conn)
	log.Info("Connected!")
	return conn, nil
}

// handleReconnect will wait for a channel error
// and then continuously attempt to re-initialize both channels
func (client *Client) handleReInit(conn *amqp.Connection) bool {
	for {
		client.IsReady = false

		err := client.init(conn)

		if err != nil {
			log.Error("Failed to initialize channel. Retrying...")

			select {
			case <-client.done:
				return true
			case <-time.After(reInitDelay):
			}
			continue
		}

		select {
		case <-client.done:
			return true
		case <-client.notifyConnClose:
			log.Info("Connection closed. Reconnecting...")
			return false
		case <-client.notifyChanClose:
			log.Info("Channel closed. Re-running init...")
		}
	}
}

// init will initialize channel & declare queue
func (client *Client) init(conn *amqp.Connection) error {
	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	err = ch.Confirm(false)

	if err != nil {
		return err
	}
	_, err = ch.QueueDeclare(
		client.queueName,
		false, // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)

	if err != nil {
		return err
	}

	client.changeChannel(ch)
	client.IsReady = true

	return nil
}

// changeConnection takes a new connection to the queue,
// and updates the close listener to reflect this.
func (client *Client) changeConnection(connection *amqp.Connection) {
	client.connection = connection
	client.notifyConnClose = make(chan *amqp.Error)
	client.connection.NotifyClose(client.notifyConnClose)
}

// changeChannel takes a new channel to the queue,
// and updates the channel listeners to reflect this.
func (client *Client) changeChannel(channel *amqp.Channel) {
	client.channel = channel
	client.notifyChanClose = make(chan *amqp.Error)
	client.notifyConfirm = make(chan amqp.Confirmation, 1)
	client.channel.NotifyClose(client.notifyChanClose)
	client.channel.NotifyPublish(client.notifyConfirm)
}

// Push will push data onto the queue, and wait for a confirm.
// If no confirms are received until within the resendTimeout,
// it continuously re-sends messages until a confirm is received.
// This will block until the server sends a confirm. Errors are
// only returned if the push action itself fails, see UnsafePush.
func (client *Client) Push(data []byte) error {
	if !client.IsReady {
		return errors.New("failed to push: not connected")
	}
	for {
		err := client.UnsafePush(data)
		if err != nil {
			log.Error("Push failed. Retrying...")
			select {
			case <-client.done:
				return errShutdown
			case <-time.After(resendDelay):
			}
			continue
		}
		select {
		case confirm := <-client.notifyConfirm:
			if confirm.Ack {
				// log.Info("Push confirmed!")
				return nil
			}
		case <-time.After(resendDelay):
		}
		log.Info("Push didn't confirm. Retrying...")
	}
}

// UnsafePush will push to the queue without checking for
// confirmation. It returns an error if it fails to connect.
// No guarantees are provided for whether the server will
// receive the message.
func (client *Client) UnsafePush(data []byte) error {
	if !client.IsReady {
		return errNotConnected
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return client.channel.PublishWithContext(
		ctx,
		"",               // Exchange
		client.queueName, // Routing key
		false,            // Mandatory
		false,            // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
}

// Consume will continuously put queue items on the channel.
// It is required to call delivery.Ack when it has been
// successfully processed, or delivery.Nack when it fails.
// Ignoring this will cause data to build up on the server.
func (client *Client) Consume(consumer string) (<-chan amqp.Delivery, error) {
	if !client.IsReady {
		return nil, errNotConnected
	}
	return client.channel.Consume(
		client.queueName,
		consumer, // Consumer
		false,    // Auto-Ack
		false,    // Exclusive
		false,    // No-local
		false,    // No-Wait
		nil,      // Args
	)
}

func (client *Client) CancelConsume(consumer string) error {
	// will close() the deliveries channel
	return client.channel.Cancel(consumer, true)
}

// Close will cleanly shutdown the channel and connection.
func (client *Client) Close() error {
	if !client.IsReady {
		return errAlreadyClosed
	}
	close(client.done)
	err := client.channel.Close()
	if err != nil {
		return err
	}
	err = client.connection.Close()
	if err != nil {
		return err
	}

	client.IsReady = false
	return nil
}
