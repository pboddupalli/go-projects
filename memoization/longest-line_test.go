package memoization

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestLongestLine(t *testing.T) {
  mat := [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}}
  // assert.Equal(t, 3, longestLine(mat))

  mat = [][]int{{1, 1, 1, 1}, {0, 1, 1, 0}, {0, 0, 0, 1}}
  assert.Equal(t, 4, longestLine(mat))
}
