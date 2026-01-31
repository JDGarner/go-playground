package cachev1_test

import (
	"testing"

	cachev1 "github.com/JDGarner/go-playground/systems/cache/v1"
	"github.com/stretchr/testify/assert"
)

type User struct {
	name string
}

func TestCache(t *testing.T) {
	t.Run("can set and get items", func(t *testing.T) {
		c := cachev1.New[User]()

		c.Set("snowbell", User{
			name: "snowy",
		})
		c.Set("chunko", User{
			name: "chunko",
		})

		chunko, ok := c.Get("chunko")
		assert.Equal(t, "chunko", chunko.name)
		assert.True(t, ok)

		snowy, ok := c.Get("snowbell")
		assert.Equal(t, "snowy", snowy.name)
		assert.True(t, ok)
	})
}
