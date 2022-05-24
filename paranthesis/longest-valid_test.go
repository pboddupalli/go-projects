package paranthesis

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestValidParan(t *testing.T) {
  assert.Equal(t, 4, longestValidParanthesis(")()())"))
}
