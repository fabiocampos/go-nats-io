package services

import (
	"github.com/fabiocampos/go-nats-io/models"
	"github.com/nats-io/go-nats"
)

type PublisherService struct {
	natsConnection *nats.Conn
	encodedConn    *nats.EncodedConn
}

// NewService creates a new service
func NewPublisherService(nc *nats.Conn, ec *nats.EncodedConn) *PublisherService {
	return &PublisherService{natsConnection: nc, encodedConn: ec}
}

//Publish string messages
func (c *PublisherService) PublishMessage(message *models.Message) {
	// Simple Publisher
	c.natsConnection.Publish("msgChannel", []byte(message.Message))
}

//Publish encoded messages
func (c *PublisherService) PublishEncodedMessage(message *models.Message) {
	// Simple Encoded Publisher
	c.encodedConn.Publish("encodedChannel", message)
}
