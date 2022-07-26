package palindromes

func palindromePairs(words []string) [][]int {

  var result [][]int
  wordMap := make(map[string]int)
  for i, word := range words {
    wordMap[word] = i
  }
  for i, word := range words {

    reversedWord := reverse(word)
    j, ok := wordMap[reversedWord]
    if ok && i != j {
      result = append(result, []int{i, j})
    }

    if isPalindrome(word) {
      j, ok = wordMap[""]
      if ok && i != j {
        result = append(result, []int{i, j}, []int{j, i})
      }
    }

    prefixes := getValidPrefixes(word) // suffixes are palindromes
    for _, prefix := range prefixes {
      r := reverse(prefix)
      j, ok := wordMap[r]
      if ok && j != i {
        result = append(result, []int{i,j})
      }
    }

    suffixes := getValidSuffixes(word) // prefixes are palindromes
    for _, suffix := range suffixes {
      r := reverse(suffix)
      j, ok := wordMap[r]
      if ok && j != i {
        result = append(result, []int{j,i})
      }
    }

  } // end of for loop

  return result
}

func isPalindrome(word string) bool {
  i, j := 0, len(word)-1
  for i <= j && word[i] == word[j] {
    i++; j--
  }
  return i > j
}

func reverse(word string) string {
  result := []byte(word)
  for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
    result[i], result[j] = result[j], result[i]
  }
  return string(result)
}

func getValidPrefixes(word string) []string {
  // suffixes are palindrome
  result := make([]string, 0)
  for i := len(word)-1; i > 0; i-- {
    if isPalindrome(word[i:]) {
      result = append(result, word[:i])
    }
  }
  return result
}

func getValidSuffixes(word string) []string {
  result := make([]string, 0)
  for i := 1; i < len(word); i++ {
    if isPalindrome(word[0:i]) {
      result = append(result, word[i:])
    }
  }
  return result
}
