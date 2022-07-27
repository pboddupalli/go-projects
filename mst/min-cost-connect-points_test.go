package mst

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMinCostConnectPoints(t *testing.T) {
  
  var points [][]int
  
  points = [][]int{{0,0},{2,2},{3,10},{5,2},{7,0}}
  assert.Equal(t, 20, minCostConnectPoints(points))

  points = [][]int{{3,12},{-2,5},{-4,1}}
  assert.Equal(t, 18, minCostConnectPoints(points))
  
  points = [][]int{{2,-3},{-17,-8},{13,8},{-17,-15}}
  assert.Equal(t, 53, minCostConnectPoints(points))

  points = [][]int{{-14,-14},{-18,5},{18,-10},{18,18},{10,-2}}
  assert.Equal(t, 102, minCostConnectPoints(points))
}
