package twod

func placeWordInCrossword(board [][]byte, word string) bool {

  m, n := len(board), len(board[0])

  rword := reverseWord(word)

  for i := 0; i < m; i++ {
    prev := -1
    gaps := [][]int{}
    for j := 0; j < n; j++ {
      if board[i][j] == '#' {
        gaps = append(gaps, []int{prev+1, j})
        prev = j
      }
    }
    gaps = append(gaps, []int{prev+1, n})

    for _, gap := range gaps {
      s, e := gap[0], gap[1]
      if e-s == len(word) {
        var k int
        for k = 0; k < len(word) && (board[i][s+k] == ' ' || board[i][s+k] == word[k]); k++ {
        }
        if k == len(word) {
          return true
        }
        for k = 0; k < len(word) && (board[i][s+k] == ' ' || board[i][s+k] == rword[k]); k++ {
        }
        if k == len(word) {
          return true
        }
      }
    }

  }

  for j := 0; j < n; j++ {
    prev := -1
    gaps := [][]int{}
    for i := 0; i < m; i++ {
      if board[i][j] == '#' {
        gaps = append(gaps, []int{prev+1, i})
        prev = i
      }
    }
    gaps = append(gaps, []int{prev+1, m})

    for _, gap := range gaps {
      s, e := gap[0], gap[1]
      if e-s == len(word) {
        var k int
        for k = 0; k < len(word) && (board[s+k][j] == ' ' || board[s+k][j] == word[k]); k++ {
        }
        if k == len(word) {
          return true
        }
        for k = 0; k < len(word) && (board[s+k][j] == ' ' || board[s+k][j] == rword[k]); k++ {
        }
        if k == len(word) {
          return true
        }
      }
    }
  }

  return false
}

func reverseWord(word string) string {
  n := len(word)
  retval := make([]byte, n)
  for i := 0; i < n; i++ {
    retval[n-1-i] = word[i]
  }
  return string(retval)
}
