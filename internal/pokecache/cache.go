package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    sync.Mutex
	items map[string]cacheEntry
}

func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newCacheEntry := cacheEntry{
		val:       data,
		createdAt: time.Now(),
	}
	c.items[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.items[key]

	if !exists {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			<-ticker.C

			c.mu.Lock()
			now := time.Now()
			for key, value := range c.items {
				if now.Sub(value.createdAt) > interval {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		items: make(map[string]cacheEntry),
	}
	newCache.reapLoop(interval)
	return newCache
}
