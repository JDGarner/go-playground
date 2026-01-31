package cachev2

import (
	"fmt"
	"testing"
)

// func TestBasicSetAndGet(t *testing.T) {
// 	cache := New[int](16)

// 	cache.Set("foo", 42)
// 	val, ok := cache.Get("foo")

// 	if !ok {
// 		t.Fatal("expected key 'foo' to be present")
// 	}
// 	if val != 42 {
// 		t.Errorf("expected value 42, got %d", val)
// 	}
// }

// func TestGetMissingKey(t *testing.T) {
// 	cache := New[int](16)

// 	val, ok := cache.Get("nonexistent")

// 	if ok {
// 		t.Error("expected ok to be false for missing key")
// 	}
// 	if val != 0 {
// 		t.Errorf("expected zero value, got %d", val)
// 	}
// }

// func TestOverwriteExistingKey(t *testing.T) {
// 	cache := New[int](16)

// 	cache.Set("foo", 1)
// 	cache.Set("foo", 2)

// 	val, ok := cache.Get("foo")
// 	if !ok {
// 		t.Fatal("expected key to be present")
// 	}
// 	if val != 2 {
// 		t.Errorf("expected value 2, got %d", val)
// 	}
// }

// func TestZeroValueStorage(t *testing.T) {
// 	cache := New[int](16)

// 	cache.Set("zero", 0)

// 	val, ok := cache.Get("zero")
// 	if !ok {
// 		t.Fatal("expected key 'zero' to be present even though value is zero")
// 	}
// 	if val != 0 {
// 		t.Errorf("expected value 0, got %d", val)
// 	}
// }

// func TestDeleteExistingKey(t *testing.T) {
// 	cache := New[int](16)

// 	cache.Set("foo", 42)
// 	deleted := cache.Delete("foo")

// 	if !deleted {
// 		t.Error("expected Delete to return true for existing key")
// 	}

// 	_, ok := cache.Get("foo")
// 	if ok {
// 		t.Error("expected key to be gone after deletion")
// 	}
// }

// func TestDeleteMissingKey(t *testing.T) {
// 	cache := New[int](16)

// 	deleted := cache.Delete("nonexistent")

// 	if deleted {
// 		t.Error("expected Delete to return false for missing key")
// 	}
// }

// func TestDeleteAlreadyDeleted(t *testing.T) {
// 	cache := New[int](16)

// 	cache.Set("foo", 42)
// 	cache.Delete("foo")
// 	deleted := cache.Delete("foo")

// 	if deleted {
// 		t.Error("expected second Delete to return false")
// 	}
// }

// // TestTombstoneProbeChain verifies that deleting an entry doesn't break
// // the probe chain for other entries that were placed after it due to collision.
// func TestTombstoneProbeChain(t *testing.T) {
// 	cache := New[string](16)

// 	// We need two keys that hash to the same index.
// 	// Find two such keys by brute force.
// 	key1, key2 := findCollidingKeys(cache)

// 	cache.Set(key1, "first")
// 	cache.Set(key2, "second")

// 	// Delete the first key (creates a tombstone)
// 	cache.Delete(key1)

// 	// The second key should still be findable
// 	val, ok := cache.Get(key2)
// 	if !ok {
// 		t.Fatalf("expected key2 '%s' to still be present after deleting key1 '%s'", key2, key1)
// 	}
// 	if val != "second" {
// 		t.Errorf("expected value 'second', got '%s'", val)
// 	}
// }

// // TestCollisionCorrectValue verifies that when two keys collide,
// // Get returns the correct value for each key, not the other's value.
// func TestCollisionCorrectValue(t *testing.T) {
// 	cache := New[string](16)

// 	key1, key2 := findCollidingKeys(cache)

// 	cache.Set(key1, "value1")
// 	cache.Set(key2, "value2")

// 	val1, ok1 := cache.Get(key1)
// 	val2, ok2 := cache.Get(key2)

// 	if !ok1 || !ok2 {
// 		t.Fatal("both keys should be present")
// 	}
// 	if val1 != "value1" {
// 		t.Errorf("key1 should have 'value1', got '%s'", val1)
// 	}
// 	if val2 != "value2" {
// 		t.Errorf("key2 should have 'value2', got '%s'", val2)
// 	}
// }

// // TestSetAfterDeleteReusesTombstone verifies that Set can reuse a tombstone slot.
// func TestSetAfterDeleteReusesTombstone(t *testing.T) {
// 	cache := New[int](16)

// 	cache.Set("foo", 1)
// 	cache.Delete("foo")
// 	cache.Set("foo", 2)

// 	val, ok := cache.Get("foo")
// 	if !ok {
// 		t.Fatal("expected key to be present after re-setting")
// 	}
// 	if val != 2 {
// 		t.Errorf("expected value 2, got %d", val)
// 	}
// }

// TestSetDoesNotCreateDuplicateAfterTombstone verifies that if a key exists
// further along the probe chain past a tombstone, Set updates that entry
// rather than creating a duplicate at the tombstone position.
func TestSetDoesNotCreateDuplicateAfterTombstone(t *testing.T) {
	cache := New[string](16)

	key1, key2 := findCollidingKeys(cache)

	// key1 goes to index N, key2 probes to N+1
	cache.Set(key1, "first")
	cache.Set(key2, "second")

	// Delete key1, creating a tombstone at index N
	cache.Delete(key1)

	// Now update key2. A buggy implementation might insert at the tombstone
	// instead of finding and updating the existing entry at N+1.
	cache.Set(key2, "updated")

	val, ok := cache.Get(key2)
	if !ok {
		t.Fatal("key2 should be present")
	}
	if val != "updated" {
		t.Errorf("expected 'updated', got '%s'", val)
	}

	// Verify there's no duplicate by checking size is correct
	// (This assumes you track size correctly)
	if cache.size != 1 {
		t.Errorf("expected size 1, got %d (possible duplicate entry)", cache.size)
	}
}

// TestManyCollisions stress tests with multiple keys hashing to same index.
// func TestManyCollisions(t *testing.T) {
// 	cache := New[int](16)

// 	// Find several keys that all hash to the same index
// 	keys := findMultipleCollidingKeys(cache, 16)

// 	for i, key := range keys {
// 		cache.Set(key, i)
// 	}

// 	// All should be retrievable with correct values
// 	for i, key := range keys {
// 		val, ok := cache.Get(key)
// 		if !ok {
// 			t.Errorf("key '%s' should be present", key)
// 		}
// 		if val != i {
// 			t.Errorf("key '%s' should have value %d, got %d", key, i, val)
// 		}
// 	}
// }

// // TestDeleteMiddleOfChain deletes an entry in the middle of a collision chain.
// func TestDeleteMiddleOfChain(t *testing.T) {
// 	cache := New[int](16)

// 	keys := findMultipleCollidingKeys(cache, 16)

// 	cache.Set(keys[0], 0)
// 	cache.Set(keys[1], 1)
// 	cache.Set(keys[2], 2)

// 	// Delete the middle one
// 	cache.Delete(keys[1])

// 	// First and last should still work
// 	val0, ok0 := cache.Get(keys[0])
// 	val2, ok2 := cache.Get(keys[2])

// 	if !ok0 || val0 != 0 {
// 		t.Errorf("keys[0] should have value 0")
// 	}
// 	if !ok2 || val2 != 2 {
// 		t.Errorf("keys[2] should have value 2")
// 	}

// 	// Middle one should be gone
// 	_, ok1 := cache.Get(keys[1])
// 	if ok1 {
// 		t.Error("keys[1] should be deleted")
// 	}
// }

// // TestSizeTracking verifies size is tracked correctly through operations.
// func TestSizeTracking(t *testing.T) {
// 	cache := New[int](16)

// 	if cache.size != 0 {
// 		t.Errorf("initial size should be 0, got %d", cache.size)
// 	}

// 	cache.Set("a", 1)
// 	if cache.size != 1 {
// 		t.Errorf("size after one Set should be 1, got %d", cache.size)
// 	}

// 	cache.Set("a", 2) // overwrite, size shouldn't change
// 	if cache.size != 1 {
// 		t.Errorf("size after overwrite should still be 1, got %d", cache.size)
// 	}

// 	cache.Set("b", 3)
// 	if cache.size != 2 {
// 		t.Errorf("size after second key should be 2, got %d", cache.size)
// 	}

// 	cache.Delete("a")
// 	if cache.size != 1 {
// 		t.Errorf("size after delete should be 1, got %d", cache.size)
// 	}

// 	cache.Delete("nonexistent")
// 	if cache.size != 1 {
// 		t.Errorf("size after deleting nonexistent should still be 1, got %d", cache.size)
// 	}
// }

// Helper function to find two keys that hash to the same index.
func findCollidingKeys(cache *Cache[string]) (string, string) {
	indices := make(map[int]string)

	for i := 0; ; i++ {
		key := fmt.Sprintf("key%d", i)
		idx := cache.indexFromKey(key)

		if existing, found := indices[idx]; found {
			return existing, key
		}
		indices[idx] = key
	}
}

// Helper function to find n keys that all hash to the same index.
func findMultipleCollidingKeys(cache *Cache[int], n int) []string {
	targetIdx := -1
	var keys []string

	for i := 0; len(keys) < n; i++ {
		key := fmt.Sprintf("key%d", i)
		idx := cache.indexFromKey(key)

		if targetIdx == -1 {
			targetIdx = idx
			keys = append(keys, key)
		} else if idx == targetIdx {
			keys = append(keys, key)
		}
	}

	return keys
}
