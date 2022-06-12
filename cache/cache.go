package cache

import (
	"sync"
)

type Cache struct {
	sync.RWMutex
	Data map[string]Order
}

type Order struct {
	Id string `json:"order_uid"`
	Data string
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

func (c *Cache) Get(key string) interface{} {
	c.RLock()
	defer c.RUnlock()
	if c.Data[key].Id != "" {
		return c.Data[key]
	}
	return nil
}