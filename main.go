package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	productionEnv := flag.Bool("production", false, "When production is activated, the log level is set to Warn")

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if *productionEnv == true {
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

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
