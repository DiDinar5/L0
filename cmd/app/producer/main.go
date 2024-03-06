package main

import (
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

const (
	ClusterID = "test-cluster"
	ClientID  = "producer"
	NatsURL   = "nats://localhost:4222"
)

func main() {
	sc, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(NatsURL))
	log.Println("Successfully connected to nats-streaming from producer")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	byteValue, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatal("Error model.josn")
	}

	sc.Publish("orders", []byte(byteValue))
}
