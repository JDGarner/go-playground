package cachev2

import (
	"hash/fnv"
)

type Cache[T any] struct {
	data     []Entry[T]
	capacity int
	size     int
}

type Entry[T any] struct {
	key       string
	value     T
	present   bool
	tombstone bool
}

func New[T any](capacity int) *Cache[T] {
	return &Cache[T]{
		data:     make([]Entry[T], capacity),
		capacity: capacity,
		size:     0,
	}
}

func (c *Cache[T]) Set(key string, value T) {
	idx := c.indexFromKey(key)

	// If the entry is not taken, then insert here
	if !c.data[idx].present {
		c.insert(idx, key, value)
		return
	}

	// Otherwise, if this entry exists already (matching key) we need to insert it there
	// Else we need to insert it at the first empty or tombstoned spot
	// (This is to make sure we override tombstoned slots ONLY if there's not a duplicate key already in the cache)
	firstTombstonedIdx := -1

	// Otherwise keep probing while the entry is present
	for c.data[idx].present {
		entry := c.data[idx]

		// if the key is the same overwrite it
		if entry.key == key {
			if entry.tombstone {
				c.insert(idx, key, value) // if it's a tombstone we want to incr the cache size
			} else {
				c.overwrite(idx, key, value)
			}

			return
		}

		if entry.tombstone && firstTombstonedIdx == -1 {
			firstTombstonedIdx = idx
		}

		// otherwise, try the next slot
		idx = (idx + 1) % c.capacity
	}

	// if we reach here, we didn't find any matching keys

	// first check to see if we can insert it at the firstTombstonedIdx
	if firstTombstonedIdx != -1 {
		c.insert(firstTombstonedIdx, key, value)
		return
	}

	// if there wasn't any earlier tombstoned slot, then insert it at the first completely empty slot (idx)
	c.insert(idx, key, value)
}

func (c *Cache[T]) overwrite(index int, key string, value T) {
	c.data[index] = Entry[T]{
		key:     key,
		value:   value,
		present: true,
	}
}

func (c *Cache[T]) insert(index int, key string, value T) {
	c.data[index] = Entry[T]{
		key:     key,
		value:   value,
		present: true,
	}
	c.size++

	if c.size > int(float64(c.capacity)*0.75) {
		c.grow()
	}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	entry, _, ok := c.getEntry(key)
	if !ok {
		var zeroValue T
		return zeroValue, false
	}

	return entry.value, true
}

func (c *Cache[T]) Delete(key string) bool {
	_, idx, ok := c.getEntry(key)
	if !ok {
		return false
	}

	c.data[idx].tombstone = true
	c.size--
	return true
}

// Finds an entry in the cache that matches the given key and hasn't been tombstoned
// If none exist, returns -1
func (c *Cache[T]) getEntry(key string) (Entry[T], int, bool) {
	idx := c.indexFromKey(key)

	// Keep looking until we find an entry where the key matches OR an empty entry
	for c.data[idx].present {
		if c.data[idx].tombstone {
			idx = (idx + 1) % c.capacity
			continue
		}

		entry := c.data[idx]

		if entry.key == key {
			return entry, idx, true
		}

		idx = (idx + 1) % c.capacity
	}

	return Entry[T]{}, -1, false
}

// Create a number between 0 and c.size-1 based on the hash of the key
func (c *Cache[T]) indexFromKey(key string) int {
	hash := fnv.New64a()
	hash.Write([]byte(key))

	return int(hash.Sum64() % uint64(c.capacity))
}

func (c *Cache[T]) grow() {
	oldData := c.data

	c.capacity = c.capacity * 2
	c.data = make([]Entry[T], c.capacity)
	c.size = 0

	for _, entry := range oldData {
		if !entry.present || entry.tombstone {
			continue
		}
		c.Set(entry.key, entry.value)
	}
}
