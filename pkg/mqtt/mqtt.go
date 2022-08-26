// Package mqtt wraps github.com/eclipse/paho.mqtt.golang
package mqtt

// Server: https://mosquitto.org/download

import (
	"time"

	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/container/gmap"
)

type MQTTMsg = mqtt.Message

// Cache all mqtt clients and subscribers
var mqttClientMap = gmap.New(true)

type MQTT struct {
	Cli      mqtt.Client
	Opt      *mqtt.ClientOptions
	Subs     map[string]byte
	Timeout  time.Duration
	Delivery chan MQTTMsg
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	optReader := client.OptionsReader()
	mqtt := findMQTTByClientId((&optReader).ClientID())
	if mqtt != nil && mqtt.Cli.IsConnected() {
		// log.Infofi("%v Received message: %s from topic: %s", (&optReader).ClientID(), msg.Payload(), msg.Topic())
		mqtt.Delivery <- msg
	}
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	optReader := client.OptionsReader()
	log.Errorfi("%v disconnected with: %v", (&optReader).ClientID(), err)
}

var onConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	optReader := client.OptionsReader()
	log.Errorfi("%v connected, subscriber restore.", (&optReader).ClientID())

	mqtt := findMQTTByClientId((&optReader).ClientID())
	if mqtt != nil {
		for topic, qos := range mqtt.Subs {
			mqtt.Cli.Subscribe(topic, qos, messagePubHandler).Wait()
		}
	}
}

var reconnectHandler mqtt.ReconnectHandler = func(_ mqtt.Client, opt *mqtt.ClientOptions) {
	log.Infofi("%v reconnecting", opt.ClientID)
}

func findMQTTByClientId(clientId string) *MQTT {
	if mqttClientMap.Contains(clientId) {
		return mqttClientMap.Get(clientId).(*MQTT)
	}
	return nil
}

// broker format: tcp://broker.emqx.io:1883
func New(broker string, username string, password string, clientId string) *MQTT {
	mqttOpts := mqtt.NewClientOptions()
	mqttOpts.AddBroker(broker)
	mqttOpts.SetClientID(clientId)
	mqttOpts.SetUsername(username)
	mqttOpts.SetPassword(password)
	// Make sure you can still receive offline messages after automatic reconnection
	mqttOpts.SetCleanSession(false) //default as true
	mqttOpts.SetResumeSubs(true)    //default as false
	mqttOpts.SetAutoReconnect(true)
	mqttOpts.SetConnectRetry(true)
	mqttOpts.SetConnectRetryInterval(config.GetDurationWithDefault("mqtt.retryInterval", 5) * time.Second)
	mqttOpts.SetConnectTimeout(config.GetDurationWithDefault("mqtt.timeout", 5) * time.Second)
	mqttOpts.SetKeepAlive(config.GetDurationWithDefault("mqtt.keepAlive", 5) * time.Second)

	mqttOpts.OnConnect = onConnectHandler
	mqttOpts.OnReconnecting = reconnectHandler
	mqttOpts.SetDefaultPublishHandler(messagePubHandler)
	mqttOpts.OnConnectionLost = connectLostHandler
	client := MQTT{
		Cli:      mqtt.NewClient(mqttOpts),
		Opt:      mqttOpts,
		Subs:     make(map[string]byte),
		Delivery: make(chan mqtt.Message),
		Timeout:  config.GetDurationWithDefault("mqtt.timeout", 5) * time.Second,
	}
	token := client.Cli.Connect()
	if !token.WaitTimeout(client.Timeout) {
		log.Errorfi("%v connect timeout", clientId)
	}
	if token.Error() != nil {
		log.Errori(token.Error())
	}
	return &client
}

func (mqtt *MQTT) Close() {
	mqtt.Cli.Disconnect(0)
	// close(mqtt.Delivery)
}

func (mqtt *MQTT) Consume() <-chan MQTTMsg {
	return mqtt.Delivery
}

func (mqtt *MQTT) Publish(topic string, qos byte, retained bool, payload interface{}) {
	if !mqtt.Cli.IsConnected() {
		log.Errorfi("%s is not connected", mqtt.Opt.ClientID)
		return
	}
	if !mqtt.Cli.Publish(topic, qos, retained, payload).WaitTimeout(mqtt.Timeout) {
		log.Errorfi("%v publish timeout", mqtt.Opt.ClientID)
	}
}

func (mqtt *MQTT) Subscribe(topic string, qos byte) {
	if !mqtt.Cli.IsConnected() {
		log.Errorfi("%s is not connected", mqtt.Opt.ClientID)
		return
	}

	if !mqttClientMap.Contains(mqtt.Opt.ClientID) {
		mqttClientMap.Set(mqtt.Opt.ClientID, mqtt)
	}

	mqtt.Subs[topic] = qos

	if !mqtt.Cli.Subscribe(topic, qos, messagePubHandler).WaitTimeout(mqtt.Timeout) {
		log.Errorfi("%v subscribe timeout", mqtt.Opt.ClientID)
	}
}

func (mqtt *MQTT) UnSubscribe(topic string) {
	if !mqtt.Cli.IsConnected() {
		log.Errorfi("%s is not connected", mqtt.Opt.ClientID)
		return
	}

	if !mqttClientMap.Contains(mqtt.Opt.ClientID) {
		mqttClientMap.Set(mqtt.Opt.ClientID, mqtt)
	}

	delete(mqtt.Subs, topic)

	if !mqtt.Cli.Unsubscribe(topic).WaitTimeout(mqtt.Timeout) {
		log.Errorfi("%v unsubscribe timeout", mqtt.Opt.ClientID)
	}
}
