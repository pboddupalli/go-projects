package strings

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestLongestPrefix(t *testing.T) {
  assert.Equal(t, "fl", longestCommonPrefix([]string{"flower","flow","flight"}))
}
