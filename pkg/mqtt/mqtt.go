// Package mqtt wraps github.com/eclipse/paho.mqtt.golang
package mqtt

// Server: https://mosquitto.org/download

import (
	"time"

	"github.com/lovelacelee/clsgo/pkg/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	Cli mqtt.Client
	Opt *mqtt.ClientOptions
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	optReader := client.OptionsReader()
	log.Infofi("%v Received message: %s from topic: %s", (&optReader).ClientID(), msg.Payload(), msg.Topic())
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	optReader := client.OptionsReader()
	log.Errorfi("%v disconnected with: %v", (&optReader).ClientID(), err)
}

var onConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	optReader := client.OptionsReader()
	log.Errorfi("%v connected", (&optReader).ClientID())
}

var reconnectHandler mqtt.ReconnectHandler = func(_ mqtt.Client, opt *mqtt.ClientOptions) {
	log.Infofi("%v reconnecting", opt.ClientID)
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
	mqttOpts.SetConnectRetryInterval(time.Second)
	mqttOpts.SetConnectTimeout(time.Second)
	mqttOpts.SetKeepAlive(time.Second * 5)

	mqttOpts.OnConnect = onConnectHandler
	mqttOpts.OnReconnecting = reconnectHandler
	mqttOpts.SetDefaultPublishHandler(messagePubHandler)
	mqttOpts.OnConnectionLost = connectLostHandler
	client := MQTT{
		Cli: mqtt.NewClient(mqttOpts),
		Opt: mqttOpts,
	}
	if token := client.Cli.Connect(); token.Wait() && token.Error() != nil {
		log.Errori(token.Error())
	}
	return &client
}

func (mqtt *MQTT) Close() {
	mqtt.Cli.Disconnect(0)

}

func (mqtt *MQTT) Publish(topic string, qos byte, retained bool, payload interface{}) {
	if !mqtt.Cli.IsConnected() {
		log.Errorf("%s is not connected", mqtt.Opt.ClientID)
		return
	}
	mqtt.Cli.Publish(topic, qos, retained, payload).Wait()
}

func (mqtt *MQTT) Subscribe(topic string, qos byte) {
	if !mqtt.Cli.IsConnected() {
		log.Errorf("%s is not connected", mqtt.Opt.ClientID)
		return
	}
	mqtt.Cli.Subscribe(topic, qos, messagePubHandler).Wait()
}
