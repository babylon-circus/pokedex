package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	records  map[string]cacheEntry
	interval time.Duration
	mu       *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {

	c := Cache{
		records:  make(map[string]cacheEntry),
		interval: interval,
		mu:       &sync.RWMutex{},
	}
	c.reapLoop()

	return c
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	go func() {
		defer ticker.Stop()
		for range ticker.C {
			c.reap()
		}
	}()

}

func (c *Cache) reap() {
	for key, record := range c.records {
		if time.Since(record.createdAt) > c.interval {
			fmt.Printf("Deleting log entry for %s\n", key)
			delete(c.records, key)
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Printf("Adding log entry for %s\n", key)
	c.mu.Lock()
	c.records[key] = cacheEntry{val, time.Now()}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Printf("Using log entry for %s\n", key)
	c.mu.Lock()
	entry, ok := c.records[key]
	c.mu.Unlock()

	if !ok {
		return nil, false
	}

	return entry.val, true

}
