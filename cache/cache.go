package cache

import (
	"sync"
)

type Cache struct {
	sync.Mutex
	Data map[string]Order
}

type Order struct {
	Id string
	Data []byte
}

func New() *Cache {
	c := &Cache{}
	c.Data := make(map[string]Order)
	return c
}

func (c *Cache) Add(o *Order) {
	c.Lock()
	defer c.Unlock()
	c.Data[o.Id] = o
}