package arrays

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestLongestWord(t *testing.T) {
  assert.Equal(t, "world", longestWord([]string{"w","wo","wor","worl","world"}))
  assert.Equal(t, "apple", longestWord([]string{"a","banana","app","appl","ap","apply","apple"}))

  words := []string{"yo","ew","fc","zrc","yodn","fcm","qm","qmo","fcmz","z","ewq","yod","ewqz","y"}
  assert.Equal(t, "yodn", longestWord(words))
}
