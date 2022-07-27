package twod

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMinAreaRect(t *testing.T) {
  var points [][]int
  
  points = [][]int{{1,1},{1,3},{3,1},{3,3},{2,2}}
  assert.Equal(t, 4, minAreaRect(points))

  points = [][]int{{1,1},{1,3},{3,1},{3,3},{4,1},{4,3}}
  assert.Equal(t, 2, minAreaRect(points))
}
