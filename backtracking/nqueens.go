package backtracking

var rows, cols []bool
var ldiags, rdiags []bool

func solveNQueens(n int) [][]string {
  var result [][]string

  rows = make([]bool, n)
  cols = make([]bool, n)

  ldiags = make([]bool, 2 * n - 1)
  rdiags = make([]bool, 2 * n - 1)

  // fmt.Println(len(ldiags), len(rdiags))

  // board is what gets into the final result
  board := make([][]byte, n)
  for i := 0; i < n; i++ {
    board[i] = make([]byte, n)
    for j := 0; j < n; j++ {
      board[i][j] = '.'
    }
  }
  traverse(0 /* current */, n /* board size */, board, &result)
  return result
}

func traverse(row, n int, board [][]byte, result *[][]string) {
  if row == n {
    // successfully populated all rows/cols with queen...add the board to the result
    var tmp []string
    for i := 0; i < n; i++ {
      tmp = append(tmp, string(board[i]))
    }
    *result = append(*result, tmp)
    return
  }

  // loop over all cols
  for col := 0; col < n; col++ {
    if ldiags[row + col] || rdiags[row + n - col - 1] || rows[row] || cols[col] {
      continue
    }
    rows[row], cols[col], ldiags[row+col], rdiags[row + n - col - 1] = true, true, true, true
    board[row][col] = 'Q'
    traverse(row + 1, n, board, result)
    rows[row], cols[col], ldiags[row+col], rdiags[row + n - col - 1] = false, false, false, false
    board[row][col] = '.'
  }
}
