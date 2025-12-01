package hashmap

import (
	"fmt"
	"hash/fnv"
	"strings"
)

// This hash map implementation uses open addressing for handling collisions:
// - If there is a collision, put the item in the next available slot

const initialSize = 10

type HashMap[T any] struct {
	data []Entry[T]
	size int // Number of entries in hashmap, not total size of data array
}

type Entry[T any] struct {
	key   string
	value T
}

func (e *Entry[T]) isEmpty() bool {
	return e.key == ""
}

func New[T any]() *HashMap[T] {
	return &HashMap[T]{
		data: make([]Entry[T], initialSize),
		size: 0,
	}
}

// If the value retrieved key does not match, keep looking until either we find it or
// empty spot is found
func (h *HashMap[T]) Get(key string) (T, bool) {
	index := h.hash(key)
	entry := h.data[index]

	for entry.key != key {
		if entry.isEmpty() {
			var zero T
			return zero, false
		}

		index = (index + 1) % len(h.data)
		entry = h.data[index]
	}

	return entry.value, true
}

// If index is taken, put the item in next available index
func (h *HashMap[T]) Insert(key string, value T) {
	index := h.hash(key)

	for !h.data[index].isEmpty() {
		if h.data[index].key == key {
			h.data[index].value = value
			return
		}

		index = (index + 1) % len(h.data)
	}

	h.data[index] = Entry[T]{
		key:   key,
		value: value,
	}
	h.size++

	if h.size > len(h.data)/2 {
		h.Grow()
	}
}

// Double the size of data, rehash everything
func (h *HashMap[T]) Grow() {
	fmt.Println(">>> growing the hashmap!")
	doubled := make([]Entry[T], len(h.data)*2)

	for _, entry := range h.data {
		if entry.isEmpty() {
			continue
		}

		index := h.hashGrow(entry.key, len(doubled))

		for doubled[index].key != "" {
			index = (index + 1) % len(doubled)
		}
		doubled[index] = entry
	}

	h.data = doubled
}

func (h *HashMap[T]) String() string {
	var sb strings.Builder

	for i, entry := range h.data {
		if !entry.isEmpty() {
			sb.WriteString(fmt.Sprintf("%s: %v (index: %d)\n", entry.key, entry.value, i))
		}
	}

	return sb.String()
}

func (h *HashMap[T]) hash(key string) int {
	fnvHash := fnv.New64a()
	fnvHash.Write([]byte(key))

	return int(fnvHash.Sum64() % (uint64(len(h.data))))
}

func (h *HashMap[T]) hashGrow(key string, newLen int) int {
	fnvHash := fnv.New64a()
	fnvHash.Write([]byte(key))

	return int(fnvHash.Sum64() % (uint64(newLen)))
}
