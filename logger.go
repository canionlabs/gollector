package main

import log "github.com/sirupsen/logrus"

func errorHandlerWithFields(err error, message string, fields map[string]interface{}) {
	if err != nil {
		log.WithFields(fields).Fatal(message)
	}
}

func errorHandler(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}
