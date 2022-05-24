package arrays

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMaxSumSubarray(t *testing.T) {
  assert.Equal(t, 6, maxSumSubarray([]int{-2,1,-3,4,-1,2,1,-5,4}))
  assert.Equal(t, 1, maxSumSubarray([]int{1}))
  assert.Equal(t, 23, maxSumSubarray([]int{5,4,-1,7,8}))
}
