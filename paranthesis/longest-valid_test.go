package paranthesis

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestValidParan(t *testing.T) {
  // assert.Equal(t, 4, longestValidParenthesesVer2(")()())"))
  // assert.Equal(t, 2, longestValidParenthesesVer2("(()"))
  // assert.Equal(t, 4, longestValidParenthesesVer3("(())("))
  assert.Equal(t, 4, longestValidParenthesesVer3(")()())()()("))
}
