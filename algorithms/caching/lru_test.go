package caching_test

import (
	"testing"

	"github.com/JDGarner/go-playground/algorithms/caching"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	t.Run("items can be added to and retrieved from the cache", func(t *testing.T) {
		cache := caching.NewLRUCache(3)

		err := cache.Add("name", "Snowbell")
		assert.NoError(t, err)

		value, err := cache.Get("name")
		assert.NoError(t, err)
		assert.Equal(t, "Snowbell", value)
	})

	t.Run("returns a cache miss error when trying to get an item that is not in the cache", func(t *testing.T) {
		cache := caching.NewLRUCache(3)

		_, err := cache.Get("something")
		assert.ErrorIs(t, err, caching.CacheMissError)
	})

	t.Run("items can be removed from the cache", func(t *testing.T) {
		cache := caching.NewLRUCache(3)

		_ = cache.Add("1", "one")
		cache.Remove("1")

		value, err := cache.Get("1")
		assert.Equal(t, "", value)
		assert.ErrorIs(t, err, caching.CacheMissError)
	})

	t.Run("if a key already exists in the cache, another item with the same key cannot be added", func(t *testing.T) {
		cache := caching.NewLRUCache(3)

		err := cache.Add("7892", "bacon")
		assert.NoError(t, err)
		err = cache.Add("7892", "tomato")
		assert.ErrorIs(t, err, caching.CacheEntryAlreadyExistsError)
	})

	t.Run("the least recently used/read item is removed from the cache when it is full", func(t *testing.T) {
		cache := caching.NewLRUCache(3)

		// Add items to the cache until it is full
		err := cache.Add("1", "one")
		assert.NoError(t, err)
		err = cache.Add("2", "two")
		assert.NoError(t, err)
		err = cache.Add("3", "three")
		assert.NoError(t, err)

		assert.Equal(t, cache.Keys(), []string{"1", "2", "3"})

		// Simulate accessing stuff from the cache in a certain order
		_, _ = cache.Get("2") // "2" is the first item we read so it will end up being the least recently used
		_, _ = cache.Get("1")
		_, _ = cache.Get("3")

		// Assert order of keys is now in order of usage recency
		assert.Equal(t, []string{"2", "1", "3"}, cache.Keys())

		// Add a new item to the cache
		err = cache.Add("4", "four")
		assert.NoError(t, err)

		// Assert least recently used item was removed, and 4 added to end of the list
		assert.Equal(t, []string{"1", "3", "4"}, cache.Keys())

		// Assert that the 'Least Recently Used/Read' item is no longer in the cache
		_, err = cache.Get("2")
		assert.ErrorIs(t, err, caching.CacheMissError)
	})
}
