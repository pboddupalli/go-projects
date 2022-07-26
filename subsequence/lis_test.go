package subsequence

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestLis(t *testing.T) {
  assert.Equal(t, 4, lis([]int{10,9,2,5,3,7,101,18}))
  assert.Equal(t, 4, lis([]int{0,1,0,3,2,3}))
  assert.Equal(t, 1, lis([]int{7,7,7}))
}
