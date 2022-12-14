package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var factsUrl = "https://catfact.ninja/fact"
var dbName = "catfact"
var collectionName = "facts"
var address = ":3000"
var fetchIntervalSeconds = 10

func main() {
	envVariables := loadEnvironmentVars()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(buildConnectionString(&envVariables)))
	if err != nil {
		panic(err)
	}

	worker := NewCatFactWorker(client, fetchIntervalSeconds)
	go worker.start()

	server := NewServer(client)

	log.Printf("INFO: Starting server on %s", address)
	http.HandleFunc("/facts", server.handleGetAllFacts)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
