package caching

import (
	"errors"
	"fmt"
)

// TODO: improvements:
// - make into a generic []T
// - change Get return to an optional?

var (
	CacheMissError               = errors.New("cache miss")
	CacheEntryAlreadyExistsError = errors.New("entry with this key already exists in cache")
)

type LRUCache struct {
	capacity int
	keys     []string // items at the end are most recently used
	itemsMap map[string]string
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		keys:     []string{},
		itemsMap: map[string]string{},
	}
}

func (c *LRUCache) Get(key string) (string, error) {
	item, ok := c.itemsMap[key]
	if ok {
		// Remove key from keys and then add it back to the 'top' (so it is considered most recently used)
		filteredKeys := []string{}
		for _, k := range c.keys {
			if k != key {
				filteredKeys = append(filteredKeys, k)
			}
		}
		filteredKeys = append(filteredKeys, key)

		c.keys = filteredKeys

		return item, nil
	}

	return "", CacheMissError
}

func (c *LRUCache) Remove(key string) {
	delete(c.itemsMap, key)

	filteredKeys := []string{}
	for _, k := range c.keys {
		if k != key {
			filteredKeys = append(filteredKeys, k)
		}
	}
	c.keys = filteredKeys
}

func (c *LRUCache) Add(key string, value string) error {
	_, exists := c.itemsMap[key]
	if exists {
		return fmt.Errorf("failed to add key '%s' to cache: %w", key, CacheEntryAlreadyExistsError)
	}

	// If the cache is already full, remove the least recently used item (key at the front of keys)
	// before adding the new entry to the cache
	if len(c.itemsMap) >= c.capacity {
		keyToRemove := c.keys[0]
		c.keys = c.keys[1:]
		delete(c.itemsMap, keyToRemove)
	}

	c.itemsMap[key] = value
	c.keys = append(c.keys, key)

	return nil
}

func (c *LRUCache) Keys() []string {
	return c.keys
}
