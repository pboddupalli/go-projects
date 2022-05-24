package palindromes

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCountSubstrings(t *testing.T) {
  assert.Equal(t, 3, countSubstrings("abc"))
  assert.Equal(t, 6, countSubstrings("aaa"))
}
