package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

type CatFactWorker struct {
	client   *mongo.Client
	interval int
}

func NewCatFactWorker(c *mongo.Client, interval int) *CatFactWorker {
	return &CatFactWorker{
		client:   c,
		interval: interval,
	}
}

func (cfw *CatFactWorker) start() error {
	coll := cfw.client.Database(dbName).Collection(collectionName)
	ticker := time.NewTicker(time.Duration(cfw.interval) * time.Second)

	for {
		resp, err := http.Get(factsUrl)
		if err != nil {
			return err
		}
		var catFact bson.M // map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}

		_, err = coll.InsertOne(context.TODO(), catFact)
		if err != nil {
			return err
		}

		log.Printf("INFO: new cat fact added")
		<-ticker.C
	}
}
