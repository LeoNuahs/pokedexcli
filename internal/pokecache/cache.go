package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	store    map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		store:    make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.store[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

// Clears all the entries that exceed the given interval
func (c *Cache) reapLoop() {
	timer := time.NewTicker(c.interval)

	for {
		<-timer.C // Wait for a tick before proceeding

		c.mu.Lock()
		for key, val := range c.store {
			if time.Since(val.createdAt) >= c.interval {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}
}
