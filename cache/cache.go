package cache

import (
	"L0/dbconnection/entity"
	"log"
	"sync"
)

type Cache struct {
	mu    sync.Mutex
	store map[string]entity.Order
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]entity.Order),
	}
}

func (c *Cache) Set(key string, value entity.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
	log.Printf("Order with ID %s was saved to cache", key)
}

func (c *Cache) Get(key string) (entity.Order, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, found := c.store[key]
	return val, found
}
