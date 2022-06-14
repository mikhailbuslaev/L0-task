package cache

import (
	"errors"
	"sync"
)

type Cache struct {
	sync.RWMutex
	Data map[string]Order
}

type Order struct {
	Id string `json:"order_uid"`
	Data string `json:"data"`
}

func New() *Cache {
	c := &Cache{}
	c.Data = make(map[string]Order)
	return c
}

func (c *Cache) Add(o *Order) {
	c.Lock()
	defer c.Unlock()
	c.Data[o.Id] = *o
}

func (c *Cache) Get(key string) (Order, error) {
	c.RLock()
	defer c.RUnlock()
	if c.Data[key].Id != "" {
		return c.Data[key], nil
	}
	return c.Data[key], errors.New("No order here")
}