package caching_test

import (
	"testing"

	"github.com/JDGarner/go-playground/algorithms/caching"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) { 
	t.Run("the least recently used/read item is removed from the cache when it is full", func(t *testing.T) {

		cache := caching.NewLRUCache(3)

		// Add items to the cache until it is full
		cache.Add("1", "one")
		cache.Add("2", "two")
		cache.Add("3", "three")

		// TODO: add GetAllKeys / AllKeys func to assert what is in the cache here

		// Get an item from the cache
		_, _ = cache.Get("2") // "2" is the first item we read so it will end up being the least recently used
		_, _ = cache.Get("1")
		_, _ = cache.Get("3")

		// Add a new item to the cache
		cache.Add("4", "four")

		// Assert that the 'Least Recently Used/Read' item is no longer in the cache
		_, err := cache.Get("2")
		assert.ErrorIs(t, err, caching.CacheMissError)

		// TODO: add GetAllKeys / AllKeys func to assert what is in the cache here
	})
}