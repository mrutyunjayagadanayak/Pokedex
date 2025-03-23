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

func NewCache(interval time.Duration) error {

	return nil
}
