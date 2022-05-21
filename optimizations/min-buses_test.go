package optimizations

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMinBuses(t *testing.T) {
  routes := [][]int{{1,2,7},{3,6,7}}
  assert.Equal(t, 2, numBusesToDestination(routes, 1, 6))

  routes = [][]int{{7,12},{4,5,15},{6},{15,19},{9,12,13}}
  assert.Equal(t, -1, numBusesToDestination(routes, 15, 12))
}
