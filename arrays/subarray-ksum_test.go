package arrays

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSubarrayKSum(t *testing.T) {
  // assert.Equal(t, true, checkSubarraySum([]int{23,2,4,6,7}, 6))
  // assert.Equal(t, true, checkSubarraySum([]int{23,2,6,4,7}, 6))
  // assert.Equal(t, false, checkSubarraySum([]int{23,2,6,4,7}, 13))
  // assert.Equal(t, false, checkSubarraySum([]int{1,0}, 2))
  assert.Equal(t, true, checkSubarraySum([]int{5,0,0,0}, 3))
}
