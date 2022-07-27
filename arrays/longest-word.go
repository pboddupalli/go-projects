package arrays

import "sort"

func longestWord(words []string) string {

  wordMap := make(map[string]struct{})
  for _, word := range words {
    wordMap[word] = struct{}{}
  }

  sort.Slice(words, func(i, j int) bool {
    if len(words[i]) == len(words[j]) {
      return words[i] > words[j]
    }
    return len(words[i]) > len(words[j])
  })

  for i := len(words)-1; i >= 0; i-- {
    word := words[i]
    _, ok := wordMap[word[:len(word)-1]]
    if ok {
      return word
    }
  }
  return ""
}
