package twod

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCrossword(t *testing.T) {
  var board [][]byte
  var word string

  board = [][]byte{{'#', ' ', '#'}, {' ', ' ', '#'}, {'#', 'c', ' '}}
  word = "abc"
  assert.Equal(t, true, placeWordInCrosswordVer2(board, word))

  board = [][]byte{{' ', '#', 'a'}, {' ', '#', 'c'}, {' ', '#', 'a'}}
  word = "ac"
  assert.Equal(t, false, placeWordInCrosswordVer2(board, word))

  board = [][]byte{{' '},{'#'},{'o'},{' '},{'t'},{'m'},{'o'},{' '},{'#'},{' '}}
  word = "octmor"
  assert.Equal(t, true, placeWordInCrosswordVer2(board, word))

  board = [][]byte{{'#',' ','#'},{'#',' ','#'},{'#',' ','c'}}
  word = "ca"
  assert.Equal(t, true, placeWordInCrosswordVer2(board, word))
}
