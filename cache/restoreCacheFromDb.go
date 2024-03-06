package cache

import (
	"L0/dbconnection/repo"
	"log"
)

func (c *Cache) RestoreCacheFromDB() {
	orders, err := repo.GetAllOrders()
	if err != nil {
		log.Printf("Failed to load orders from DB: %v", err)
		return
	}

	for _, order := range orders {
		c.Set(order.OrderUid, order)
	}

	log.Println("Cache has been restored from the database")
}
