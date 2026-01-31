package cachev1

type Cache[T any] struct {
	data map[string]T
}

func New[T any]() *Cache[T] {
	return &Cache[T]{
		data: map[string]T{},
	}
}

func (c *Cache[T]) Set(key string, value T) {
	c.data[key] = value
}

func (c *Cache[T]) Get(key string) (T, bool) {
	val, ok := c.data[key]

	return val, ok
}
