package twod

import (
  "github.com/stretchr/testify/assert"
  "testing"
)


func TestLongestPath(t *testing.T) {
  matrix := [][]int{{9,9,4},{6,6,8},{2,1,1}}
  assert.Equal(t, 4, longestIncreasingPath(matrix))
  
  matrix = [][]int{{3,4,5},{3,2,6},{2,2,1}}
  assert.Equal(t, 4, longestIncreasingPath(matrix))

  matrix = [][]int{{1}}
  assert.Equal(t, 1, longestIncreasingPath(matrix))
}