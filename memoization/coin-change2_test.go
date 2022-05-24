package memoization

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestNumCoinCombinations(t *testing.T) {
  assert.Equal(t, 4, change(5, []int{1,2,5}))
}
