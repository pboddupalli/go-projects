package strings

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestLongestWord(t *testing.T) {
  var words []string
  words = []string{"w","wo","wor","worl","world"}
  assert.Equal(t, "world", longestWord(words))

  words = []string{"a","banana","app","appl","ap","apply","apple"}
  assert.Equal(t, "apple", longestWord(words))

  words = []string{"yo","ew","fc","zrc","yodn","fcm","qm","qmo","fcmz","z","ewq","yod","ewqz","y"}
  assert.Equal(t, "yodn", longestWord(words))
}
