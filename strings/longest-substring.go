package strings

func lengthOfLongestSubstring(s string) int {
  if len(s) == 0 {
    return 0
  }
  start, maxLen := 0, 0
  var prevIndex [256]int
  for i := 0; i < 256; i++ {
    prevIndex[i] = -1
  }
  prevIndex[s[0]] = 0
  for i := 1; i < len(s); i++ {
    if prevIndex[s[i]] >= start {
      if i - start > maxLen {
        maxLen = i - start
      }
      start = prevIndex[s[i]] + 1
    }
    prevIndex[s[i]] = i
  }
  if len(s) - start > maxLen {
    maxLen = len(s) - start
  }
  return maxLen
}
