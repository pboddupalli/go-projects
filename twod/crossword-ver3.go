package twod

func placeWordInCrossword(board [][]byte, word string) bool {
  in := func(board [][]byte, i, j int) bool {
    if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
      return true
    }
    return false
  }

  for i := 0; i < len(board); i++ {
    for j := 0; j < len(board[0]); j++ {
      if board[i][j] == ' ' || board[i][j] == word[0] {

        for k := 0; k < 4; k++ { //<--- directions
          nx, ny := i-DIRS[k], j-DIRS[k+1]
          //if previous cell is on board, it must be in a '#'
          //let go for out of board
          if in(board, nx, ny) && board[nx][ny] != '#' {
            continue
          }

          //set first letter position
          nx, ny = i, j
          // place the word with the current direction k
          idx := 0 //start from the 1st char of the word
          for idx < len(word) && in(board, nx, ny) {
            if board[nx][ny] == '#' || board[nx][ny] != ' ' && board[nx][ny] != word[idx] {
              break
            }
            idx++
            nx, ny = nx+DIRS[k], ny+DIRS[k+1]
          }

          if idx == len(word) && (!in(board, nx, ny) || board[nx][ny] == '#') {
            return true
          }
        }
      } //
    } //for
  }
  return false
}

var DIRS = [...]int{0, 1, 0, -1, 0}
