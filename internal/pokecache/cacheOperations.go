package pokecache

import "time"

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
