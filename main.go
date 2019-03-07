package main

import (
	"flag"
)

func main() {
	amqpURL := flag.String("amqpUrl", "amqp://guest:guest@localhost:5672/", "AQMP URI format")
	queueName := flag.String("queue", "gollector", "AMQP Queue Name")

	mongoURL := flag.String("mongoUrl", "mongodb://localhost:27017/", "Mongo URL")
	databaseName := flag.String("database", "gollector", "Database name")
	collectionName := flag.String("collection", "gollector", "Database Collection name")

	flag.Parse()

	mongoConfig := MongoConfig{*mongoURL, *databaseName, *collectionName}
	mongo := mongoConfig.connect()

	consumer := Consumer{*amqpURL, *queueName}
	consumer.start(mongo.callbackEvent)
}
