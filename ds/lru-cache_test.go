package ds

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestLruCache(t *testing.T) {
  lru := Constructor(2)
  lru.Put(1, 1)
  lru.Put(2, 2)
  assert.Equal(t, 1, lru.Get(1))
  // lru.walk()
  // lru.walkBack()

  lru.Put(3, 3)
  assert.Equal(t, -1, lru.Get(2))
  lru.Put(4, 4)
  assert.Equal(t, -1, lru.Get(1))
  assert.Equal(t, 3, lru.Get(3))
  assert.Equal(t, 4, lru.Get(4))
}

func TestLruCacheOne(t *testing.T) {
  lru := Constructor(2)
  assert.Equal(t, -1, lru.Get(2))
  lru.Put(2,6)
  assert.Equal(t, -1, lru.Get(1))
  lru.Put(1,5)
  lru.Put(1,2)
  assert.Equal(t, 2, lru.Get(1))
  assert.Equal(t, 6, lru.Get(2))
}
