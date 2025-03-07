package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("deleting the last item from the cache", func(t *testing.T) {
		cache := NewCache(3)

		cache.Set("a", 1)
		cache.Set("b", 2)
		cache.Set("c", 3)
		cache.Set("d", 4)

		val, ok := cache.Get("a")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = cache.Get("b")
		require.True(t, ok)
		require.Equal(t, 2, val)

		val, ok = cache.Get("c")
		require.True(t, ok)
		require.Equal(t, 3, val)

		val, ok = cache.Get("d")
		require.True(t, ok)
		require.Equal(t, 4, val)
	})

	t.Run("deleting the last one used", func(t *testing.T) {
		cache := NewCache(3)

		cache.Set("a", 1)
		cache.Set("b", 2)
		cache.Set("c", 3)

		cache.Get("a")
		cache.Get("b")

		cache.Set("d", 4)

		val, ok := cache.Get("c")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
