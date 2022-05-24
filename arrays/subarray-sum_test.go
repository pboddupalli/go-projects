package arrays

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSubarraySum(t *testing.T) {
  assert.Equal(t, 2, subArraySum([]int{1,1,1}, 2))
  assert.Equal(t, 2, subArraySum([]int{1,2,3}, 3))
}
