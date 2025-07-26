package caching

import "errors"

// TODO: improvements:
// - make into a generic []T
// - change Get return to an optional?

var CacheMissError = errors.New("cache miss")

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
		// TODO: remove from keys and add back to keys at the top

		return item, nil
	}

	// TODO: client must Add this to the cache after getting a cache miss error?

	return "", CacheMissError
}

func (c *LRUCache) Remove(key string) {
	delete(c.itemsMap, key)
	// TODO: remove from keys array
}

func (c *LRUCache) Add(key string, value string) {
	if len(c.itemsMap) >= c.capacity {
		// Remove least recently used item
		// TODO: remove keys[0] from keys and delete it from itemsMap
	}

	c.itemsMap[key] = value
	c.keys = append(c.keys, key)
}
