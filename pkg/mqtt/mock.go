package mqtt

import (
	"server-poc/pkg/testutils"
	"testing"

	"github.com/matryer/is"
)

// Only use this for testing.
func Mock(t *testing.T) (Client, testutils.MQTTBroker) {
	is := is.New(t)

	// Start a broker
	broker, brokerAddr, err := testutils.StartMQTTBroker()
	is.NoErr(err) // failed to start mqtt broker

	// Connect to the broker
	client, err := Connect(brokerAddr, "", "")
	is.NoErr(err) // failed to connect to the internal broker

	return client, broker
}
