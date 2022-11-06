package MyCache

import (
	"MyCache/lru"
	"sync"
)

type cache struct {
	rw       sync.RWMutex
	lru      *lru.Cache
	maxBytes uint64
}

func (c *cache) add(key string, value ByteValue) error {
	if c.lru == nil {
		c.lru = lru.NewCache(c.maxBytes, nil)
	}
	c.rw.Lock()
	defer c.rw.Unlock()
	err := c.lru.Add(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (c *cache) remove(key string) {
	if c.lru == nil {
		c.lru = lru.NewCache(c.maxBytes, nil)
	}
	c.rw.Lock()
	defer c.rw.Unlock()
	c.lru.Remove(key)
}

func (c *cache) get(key string) (ByteValue, bool) {
	if c.lru == nil {
		c.lru = lru.NewCache(c.maxBytes, nil)
	}
	c.rw.RLock()
	defer c.rw.RUnlock()
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteValue), ok
	}
	return ByteValue{}, false
}
