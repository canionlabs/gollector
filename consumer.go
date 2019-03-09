package main

import (
	log "github.com/sirupsen/logrus"
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
	conn, err := amqp.Dial(consumer.URI)
	errorHandler(err, "Error in amqp.Dial")
	defer conn.Close()

	ch, err := conn.Channel()
	errorHandler(err, "Error in conn.Channel")
	defer ch.Close()

	queue, err := ch.QueueDeclarePassive(
		consumer.Queue, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	errorHandlerWithFields(
		err, "Error in QueueDeclare", log.Fields{"queue": consumer.Queue},
	)

	messages, _ := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	errorHandlerWithFields(
		err, "Error in Consume", log.Fields{"queue": consumer.Queue},
	)

	forever := make(chan bool)

	go func() {
		for deliveredMessage := range messages {
			callback(deliveredMessage)
			log.WithFields(log.Fields{
				"file":   "consumer",
				"queue":  queue.Name,
				"author": deliveredMessage.UserId,
			}).Info("New message received")
		}
	}()

	log.WithFields(log.Fields{
		"file":  "consumer",
		"queue": queue.Name,
	}).Warn("Start Consumer")
	<-forever
}
