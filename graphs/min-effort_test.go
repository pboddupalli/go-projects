package graphs

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMinEffort(t *testing.T) {

  var heights [][]int

  heights = [][]int{{1,2,2},{3,8,2},{5,3,5}}
  assert.Equal(t, 2, minimumEffortPath(heights))
  
  heights = [][]int{{1,2,3},{3,8,4},{5,3,5}}
  assert.Equal(t, 1, minimumEffortPath(heights))

  heights = [][]int{{1,2,1,1,1},{1,2,1,2,1},{1,2,1,2,1},{1,2,1,2,1},{1,1,1,2,1}}
  assert.Equal(t, 0, minimumEffortPath(heights))
}
