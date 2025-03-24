package pokecache

import (
	"time"
)

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
