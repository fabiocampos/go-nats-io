package services

import (
	"fmt"

	"github.com/fabiocampos/go-nats-io/models"
	"github.com/nats-io/go-nats"
)

type ConsumerService struct {
	natsConnection *nats.Conn
	encodedConn    *nats.EncodedConn
}

// NewService creates a new service
func NewConsumerService(nc *nats.Conn, ec *nats.EncodedConn) *ConsumerService {
	return &ConsumerService{natsConnection: nc, encodedConn: ec}
}

//Receives string messages
func (c *ConsumerService) ConsumeAsyncMessages(stopChannel chan int) {
	// Simple Async Subscriber
	sub, err := c.natsConnection.Subscribe("msgChannel", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	if err != nil {
		sub.Unsubscribe()
		stopChannel <- 1
	}
}

//Receives objects
func (c *ConsumerService) ConsumeEncodedMessages(stopChannel chan int) {
	// Simple Async Typed Subscriber
	sub, err := c.encodedConn.Subscribe("encodedChannel", func(m *models.Message) {
		fmt.Printf("Received an encoded message: %s\n", m.Message)
	})

	if err != nil {
		sub.Unsubscribe()
		stopChannel <- 1
	}
}
