package main

import (
	"fmt"

	"github.com/fabiocampos/go-nats-io/models"
	"github.com/fabiocampos/go-nats-io/services"
	"github.com/nats-io/go-nats"
)

func main() {
	// Connect to a nats server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Printf("Error connecting to nats-io: %s\n", err.Error())
	}
	// Starts an encoded conn to transfer json objects
	encodedConnection, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	consumerService := services.NewConsumerService(nc, encodedConnection)
	publisherService := services.NewPublisherService(nc, encodedConnection)
	stopChannel := make(chan int)

	go publisherService.PublishMessage(&models.Message{Message: "Hello World!!!"})
	go publisherService.PublishMessage(&models.Message{Message: "Is there anybody out there?"})
	go publisherService.PublishEncodedMessage(&models.Message{Message: "Now we will work of objects"})
	go publisherService.PublishEncodedMessage(&models.Message{Message: "Yes, it is really easy"})

	consumerService.ConsumeAsyncMessages(stopChannel)
	consumerService.ConsumeEncodedMessages(stopChannel)
	select {
	case stop := <-stopChannel:
		fmt.Println("Time to rest", stop)
	}
}
