package palindromes

func countSubstrings(s string) int {
  cnt := 0
  for i := 0; i < len(s); i++ {
    j, k := i, i
    for j >= 0 && k < len(s) && s[j] == s[k] {
      cnt++
      j--; k++
    }
    j, k = i, i+1
    for j >= 0 && k < len(s) && s[j] == s[k] {
      cnt++
      j--; k++
    }
  }
  return cnt
}
