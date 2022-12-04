package mqtt

import (
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/matryer/is"
	mochi "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"
)

type Broker interface {
	Publish(topic string, payload []byte, retain bool) error
	io.Closer
}

// Only use this for testing.
func Mock(t *testing.T) (Client, Broker) {
	is := is.New(t)

	// Find an open port
	port, err := findOpenPort()
	is.NoErr(err) // failed to find an open port

	// Start a broker
	broker, err := MockBroker(port)
	is.NoErr(err) // failed to start mqtt broker

	// Connect to the broker
	connectURL := fmt.Sprintf("tcp://localhost:%d", port)
	client, err := Connect(connectURL, "", "")
	is.NoErr(err) // failed to connect to the internal broker

	return client, broker
}

// Find an open port on the host machine.
func findOpenPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", ":0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()
	return port, nil
}

// Starts a broker used for testing.
func MockBroker(port int) (*mochi.Server, error) {
	// Create the broker
	broker := mochi.New()

	// Create thre tcp listener
	addr := fmt.Sprintf(":%d", port)
	tcp := listeners.NewTCP("t1", addr)

	// Add the tcp listener
	config := listeners.Config{Auth: new(auth.Allow)}
	if err := broker.AddListener(tcp, &config); err != nil {
		return nil, fmt.Errorf("failed to add a tcp listener: %w", err)
	}

	// Start the broker
	if err := broker.Serve(); err != nil {
		return nil, fmt.Errorf("failed to start serving: %w", err)
	}

	return broker, nil
}
