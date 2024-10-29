package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: make(map[string]CacheEntry),
		mu: &sync.Mutex{},
	}	
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		val: val,
		createdAt: time.Now().UTC(),
	}		
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEnt, ok := c.cache[key]
	return cacheEnt.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	timeLimit := time.Now().UTC().Add(-interval)
	for k,v := range c.cache {
		if v.createdAt.Before(timeLimit) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C{ 
		c.reap(interval);
	}
}