package palindromes

import (
  "fmt"
  "testing"
)

func TestPalindromePairs(t *testing.T) {

  var words []string

  words = []string{"abcd","dcba","lls","s","sssll"}
  fmt.Println(palindromePairs(words))

  words = []string{"bat","tab","cat"}
  fmt.Println(palindromePairs(words))

  words = []string{"a",""}
  fmt.Println(palindromePairs(words))
}
