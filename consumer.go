package main

import (
	"log"

	"github.com/streadway/amqp"
)

// Consumer provides a mapping to acess a AMQP Broker (default: RabbitMQ)
type Consumer struct {
	URI   string
	Queue string
}

// EventHandler is a callback type to be executed when
// the consumer receives new messages
type EventHandler func(amqp.Delivery)

func (consumer *Consumer) start(callback EventHandler) {
	conn, _ := amqp.Dial(consumer.URI)
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	queue, _ := ch.QueueDeclare(
		consumer.Queue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	msgs, _ := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	forever := make(chan bool)

	go func() {
		for deliveredMessage := range msgs {
			callback(deliveredMessage)
			log.Printf("Received a message: %s", deliveredMessage.Body)
		}
	}()

	<-forever
}
