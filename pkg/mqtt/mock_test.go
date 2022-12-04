package mqtt_test

import (
	"server-poc/pkg/mqtt"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestClose(t *testing.T) {
	is := is.New(t)
	mqttClient, mqttBroker := mqtt.Mock(t)
	is.NoErr(mqttClient.Close()) // Client should close correctly
	is.NoErr(mqttBroker.Close()) // Broker should close correctly
}

func TestSubscribe(t *testing.T) {
	// Setup
	is := is.New(t)
	mqttClient, mqttBroker := mqtt.Mock(t)
	defer mqttClient.Close()
	defer mqttBroker.Close()

	// Subscribe to a topic
	subData := make(chan []byte, 1)
	err := mqttClient.Sub("a/b/c", 1, func(b []byte) {
		subData <- b
	})
	is.NoErr(err) // Should subscribe to a topic

	// Publish to the topic
	payload := []byte("hello")
	err = mqttBroker.Publish("a/b/c", payload, false)
	is.NoErr(err) // Should publish on a topic

	// Wait for notification from the subscribtion
	var dataReceived []byte
	select {
	case <-time.After(time.Millisecond * 10):
		is.Fail() // Timeout occurred
	case dataReceived = <-subData:
	}

	// Make sure data stayed in tact
	is.Equal(payload, dataReceived)
}
