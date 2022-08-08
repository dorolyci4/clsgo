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

// Params for ExchangeDeclare
type Exchange struct {
	Durable      bool
	ExchangeName string
	ExchangeType string // The common types are "direct", "fanout", "topic" and "headers".
	Internal     bool
	AutoDelete   bool
	Nowait       bool
}

type Queue struct {
	QueueName  string
	Exclusive  bool
	Durable    bool
	AutoDelete bool
	Nowait     bool
}

// Rabbit MQ client
type Client struct {
	host            string
	routingKey      string
	consumerTag     string
	exchange        Exchange
	queue           Queue
	connection      *amqp.Connection
	channel         *amqp.Channel
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	// When client closed, <-done will receive false
	done chan bool
	// When IsReady is true, everything is ok for push and consume
	IsReady bool
	// Notify client is ready
	Connected chan bool
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
	errNotConnected    = errors.New("not connected to a server")
	errAlreadyClosed   = errors.New("already closed: not connected to the server")
	errShutdown        = errors.New("client is shutting down")
	errExchangeDeclare = errors.New("exchange declare failed")
	errQueueDeclare    = errors.New("queue declare failed")
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
	log.Infofi("MQ reconnect:%v reinit:%v resend:%v", reconnect, reinit, resend)
}

// New creates a new consumer state instance, and automatically
// attempts to connect to the server.
func New(addr string, exchange Exchange, queue Queue, routingKey string, consumerTag string) *Client {
	if !strings.Contains(addr, "@") {
		log.Errori("MQ address invalid, use right formmat: amqp://test:test@localhost:5672/")
		return nil
	}
	client := Client{
		host:        strings.SplitAfter(addr, "@")[1],
		queue:       queue,
		exchange:    exchange,
		routingKey:  routingKey,
		consumerTag: consumerTag,
		done:        make(chan bool),
		Connected:   make(chan bool),
	}
	go client.handleReconnect(addr)
	return &client
}

// handleReconnect will wait for a connection error on
// notifyConnClose, and then continuously attempt to reconnect.
func (client *Client) handleReconnect(addr string) {
	for {
		client.IsReady = false
		log.Infoi("AMQP attempting to connect: ", client.host)

		conn, err := client.connect(addr)

		if err != nil {
			log.Errori("Failed to connect. Retrying...")

			select {
			case <-client.done:
				return
			case <-time.After(reconnectDelay):
			}
			continue
		}

		if initialized := client.handleReInit(conn); initialized {
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
	log.Infoi("Connected!")
	return conn, nil
}

// handleReconnect will wait for a channel error
// and then continuously attempt to re-initialize both channels
func (client *Client) handleReInit(conn *amqp.Connection) bool {
	for {
		client.IsReady = false

		err := client.init(conn)

		if err != nil {
			log.Errori("Failed to initialize channel. Retrying...")

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
			log.Infoi("Connection closed. Reconnecting...")
			return false
		case <-client.notifyChanClose:
			log.Infoi("Channel closed. Re-running init...")
		}
	}
}

// init will initialize channel & declare exchange and queue
// then bind exchange and queue
func (client *Client) init(conn *amqp.Connection) error {
	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	// When noWait is true, the client will not wait for a response.
	// A channel exception could occur if the server does not support this method.
	// When ch.Confirm(false), the client will wait for a response.
	err = ch.Confirm(false)
	if err != nil {
		return err
	}
	// Set client channel first
	client.changeChannel(ch)

	// Exchange declare
	if err = client.channel.ExchangeDeclare(
		client.exchange.ExchangeName, // name of the exchange
		client.exchange.ExchangeType, // type
		client.exchange.Durable,      // durable
		client.exchange.AutoDelete,   // delete when complete
		client.exchange.Internal,     // internal
		client.exchange.Nowait,       // noWait
		nil,                          // arguments
	); err != nil {
		return errExchangeDeclare
	}

	queue, err := client.channel.QueueDeclare(
		client.queue.QueueName,
		client.queue.Durable,    // Durable
		client.queue.AutoDelete, // AutoDelete when unused
		client.queue.Exclusive,  // Exclusive
		client.queue.Nowait,     // No-wait
		nil,                     // Arguments
	)

	if err != nil {
		return errQueueDeclare
	}
	// Queue bind exchange
	if err = client.channel.QueueBind(
		queue.Name,                   // name of the queue
		client.routingKey,            // bindingKey
		client.exchange.ExchangeName, // sourceExchange
		client.queue.Nowait,          // noWait
		nil,                          // arguments
	); err != nil {
		return err
	}
	client.IsReady = true
	client.Connected <- client.IsReady

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
			log.Errori("Push failed. Retrying...")
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
				// log.Infoi("Push confirmed!")
				return nil
			}
		case <-time.After(resendDelay):
		}
		log.Infoi("Push didn't confirm. Retrying...")
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
		client.exchange.ExchangeName, // Exchange
		client.routingKey,            // Routing key
		false,                        // Mandatory
		false,                        // Immediate
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
func (client *Client) Consume(autoAck bool) (<-chan amqp.Delivery, error) {
	if !client.IsReady {
		return nil, errNotConnected
	}

	return client.channel.Consume(
		client.queue.QueueName,
		client.consumerTag,     // Consumer
		autoAck,                // Auto-Ack
		client.queue.Exclusive, // Exclusive
		false,                  // No-local, not supported by RabbitMQ
		client.queue.Nowait,    // No-Wait
		nil,                    // Args
	)
}

func (client *Client) CancelConsume() error {
	// will close() the deliveries channel
	return client.channel.Cancel(client.consumerTag, true)
}

// Close will cleanly shutdown the channel and connection.
func (client *Client) Close() error {
	if !client.IsReady {
		return errAlreadyClosed
	}
	client.IsReady = false
	close(client.done)
	close(client.Connected)
	err := client.channel.Close()
	if err != nil {
		return err
	}
	err = client.connection.Close()
	if err != nil {
		return err
	}

	return nil
}
