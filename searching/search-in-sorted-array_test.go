package searching

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSearchInSortedArray(t *testing.T) {
  assert.Equal(t, 4, srchInSortedArray([]int{4,5,6,7,0,1,2}, 0))
  assert.Equal(t, -1, srchInSortedArray([]int{4,5,6,7,0,1,2}, 3))
  assert.Equal(t, -1, srchInSortedArray([]int{1}, 0))
}

func TestSearchSortedArray(t *testing.T) {
  assert.Equal(t, true, searchVer2([]int{2,5,6,0,0,1,2}, 0))
  assert.Equal(t, false, searchVer2([]int{2,5,6,0,0,1,2}, 3))
}