package counting

import "fmt"

func bullsAndCows(secret string, guess string) string {
  secretCount := make([]int, 10)
  guessCount := make([]int, 10)
  numBulls, numCows := 0, 0
  for i := 0; i < len(secret); i++ {
    if secret[i] == guess[i] {
      numBulls++
    } else {
      guessCount[guess[i] - '0']++
      secretCount[secret[i] - '0']++
    }
  }
  for i := 0; i < len(secretCount); i++ {
    numCows += min(secretCount[i], guessCount[i])
  }
  return fmt.Sprintf("%dA%dB", numBulls, numCows)
}

func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}
