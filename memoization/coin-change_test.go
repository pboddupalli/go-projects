package memoization

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCoinChange(t *testing.T) {
  // assert.Equal(t, 3, coinChange([]int{1,2,5}, 11))
  // assert.Equal(t, -1, coinChange([]int{2}, 3))
  // assert.Equal(t, 26, coinChange([]int{216,94,15,86}, 5372))
  assert.Equal(t, 24, coinChange([]int{411,412,413,414,415,416,417,418,419,420,421,422}, 9864))
}
