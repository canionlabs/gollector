package main

import (
	"encoding/json"
	"time"

	"github.com/streadway/amqp"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MongoConfig provides a mapping to acess a Mongo database
type MongoConfig struct {
	URI          string
	DatabaseName string
	Collection   string
}

// Event provides a mapping to struct received messages
type Event struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Timestamp time.Time
	DeviceID  string
	Body      string
}

// MessageBody provides a mapping to acess the body of received messages
type MessageBody struct {
	Timestamp int64
	Body      string
}

// Collection embedding types
type Collection struct {
	*mgo.Collection
}

func (config *MongoConfig) _connect() *mgo.Collection {
	session, err := mgo.Dial(config.URI)
	errorHandler(err, "Error in mgo.Dial")
	return session.DB(config.DatabaseName).C(config.Collection)
}

func (config *MongoConfig) connect() *Collection {
	return &Collection{config._connect()}
}

func (event *Event) parseBody(body []byte) {
	msgBody := MessageBody{}
	json.Unmarshal(body, &msgBody)
	event.Body = msgBody.Body
	event.Timestamp = time.Unix(msgBody.Timestamp, 0)
}

func (mongoCollection *Collection) callbackEvent(msg amqp.Delivery) {
	event := Event{DeviceID: msg.UserId}
	event.parseBody(msg.Body)
	mongoCollection.Insert(event)
}
