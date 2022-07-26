package strings

func longestCommonPrefix(strs []string) string {
  n := len(strs)
  i := 0
  for ; i < len(strs[0]); i++ {
    j := 1
    for ; j < n && i < len(strs[j]) && strs[j][i] == strs[0][i]; j++ {
    }
    if j != n {
      break
    }
  }
  return strs[0][:i]
}
