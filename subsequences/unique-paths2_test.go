package subsequences

import (
  "github.com/stretchr/testify/assert"
  "testing"
)


func TestUniquePath2(t *testing.T) {
  grid := [][]int{{0,0,0},{0,1,0},{0,0,0}}
  assert.Equal(t, 2, uniquePaths2(grid))
}