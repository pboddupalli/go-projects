package hard

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestRussianDolls(t *testing.T) {
  assert.Equal(t, 3, russianDolls([][]int{{5,4},{6,4},{6,7},{2,3}}))
  assert.Equal(t, 1, russianDolls([][]int{{1,1},{1,1},{1,1}}))
  assert.Equal(t, 4, russianDolls([][]int{{4,5},{4,6},{6,7},{2,3},{1,1}}))
}
