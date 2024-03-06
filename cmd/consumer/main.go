package main

import (
	"L0/dbconnection/entity"
	"L0/dbconnection/repo"
	"L0/routes"

	"L0/cache"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	stan "github.com/nats-io/stan.go"
)

const (
	ClusterID = "test-cluster"
	ClientID  = "consumer"
	NatsURL   = "nats://localhost:4222"
)

func main() {
	myCache := cache.NewCache()
	if repo.HasOrdersInDB() {
		myCache.RestoreCacheFromDB()
		log.Println("Cache saved successful")
	} else {
		log.Println("No data")
	}

	sc, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(NatsURL))
	log.Println("Successfully connected to nats-streaming from consumer")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	var order entity.Order

	sub, err := sc.Subscribe("orders", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))

		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}

		id_string, err := repo.InsertOrder(order)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id_string)
		myCache.Set(order.OrderUid, order)
	})
	gin.SetMode(gin.ReleaseMode)

	router := routes.SetupRouter(myCache)
	router.Run(":8080")
	if err != nil {
		log.Fatal("error")
	}
	log.Println("This is consumer")

	time.Sleep(10 * time.Minute)
	sub.Unsubscribe()
}
