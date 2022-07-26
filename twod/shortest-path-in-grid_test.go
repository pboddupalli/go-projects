package twod

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestShortestPathInGrid(t *testing.T) {
 grid := [][]int{{0,0,0},{1,1,0},{0,0,0},{0,1,1},{0,0,0}}
 assert.Equal(t, 6, shortestPath(grid, 1))

 grid = [][]int{{0,1,1},{1,1,1},{1,0,0}}
 assert.Equal(t, -1, shortestPath(grid, 1))
}
