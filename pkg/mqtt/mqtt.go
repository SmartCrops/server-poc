package mqtt

import (
	"fmt"

	paho "github.com/eclipse/paho.mqtt.golang"
)

const (
	brokerUrl = "tcp://172.111.242.63:6666"
	username  = "roslina"
	password  = "smartcrops"
)

type Msg = paho.Message     // Msg received from a topic
type CallbackFunc func(Msg) // Function called after a message is received

// Connection to a mqtt broker
var client paho.Client

// Initialize the mqtt client
func Init() error {
	// Set client options
	opts := paho.NewClientOptions()
	opts.AddBroker(brokerUrl)
	opts.Username = username
	opts.Password = password

	// Create a client
	client = paho.NewClient(opts)

	// Connect
	t := client.Connect()
	t.Wait()
	if t.Error() != nil {
		return t.Error()
	}
	return nil
}

// Subscribe to a topic
func Subscribe(topic string, qos byte, cb CallbackFunc) error {
	t := client.Subscribe(topic, qos, func(c paho.Client, msg paho.Message) {
		cb(msg)
	})
	t.Wait()
	if t.Error() != nil {
		return fmt.Errorf("failed to subscribe to mqtt topic: %w", t.Error())
	}
	return nil
}

// Publish data on a topic
func Publish(topic string, qos byte, retained bool, data interface{}) error {
	t := client.Publish(topic, qos, retained, data)
	t.Wait()
	if t.Error() != nil {
		return fmt.Errorf("failed to send mqtt data: %w", t.Error())
	}
	return nil
}
