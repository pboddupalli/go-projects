package twod

func placeWordInCrosswordVer2(board [][]byte, word string) bool {
  m, n := len(board), len(board[0])

  for i := 0; i < m; i++ {
    k, fwd, rev := 0, true, true
    for j := 0; j < n; j++ {
      if board[i][j] == '#' {
        if k == len(word) && (fwd || rev) {
          return true
        }
        k, fwd, rev = -1, true, true
      } else if k < len(word) && board[i][j] != ' ' {
        if word[k] != board[i][j] { fwd = false }
        if word[len(word)-1-k] != board[i][j] { rev = false}
      }
      k++
    } // end of second for loop
    if k == len(word) && (fwd || rev) {
      return true
    }
  }

  for j := 0; j < n; j++ {
    k, fwd, rev := 0, true, true
    for i := 0; i < m; i++ {
      if board[i][j] == '#' {
        if k == len(word) && (fwd || rev) {
          return true
        }
        k, fwd, rev = -1, true, true
      } else if k < len(word) && board[i][j] != ' ' {
        if word[k] != board[i][j] { fwd = false }
        if word[len(word)-1-k] != board[i][j] { rev = false}
      }
      k++
    } // end of second for loop
    if k == len(word) && (fwd || rev) {
      return true
    }
  }

  return false
}
