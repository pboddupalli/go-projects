package counting

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestBullsAndCows(t *testing.T) {
  assert.Equal(t, "1A3B", bullsAndCows("1807", "7810"))
  assert.Equal(t, "1A1B", bullsAndCows("1123", "0111"))
}
