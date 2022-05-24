package numeric

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGetNumDigits(t *testing.T) {
  assert.Equal(t, 1, getNumDigits(1))
  assert.Equal(t, 4, getNumDigits(1234))
  assert.Equal(t, 5, getNumDigits(12345))
}

func TestSequentialDigits(t *testing.T) {
  fmt.Println(sequentialDigits(1000, 13000))
}