package searching

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestBinSrch(t *testing.T) {
  assert.Equal(t, 3, binSrch([]int{2,3,4,5,6}, 5))
  assert.Equal(t, -1, binSrch([]int{2,3,4,5,6}, 10))
}
