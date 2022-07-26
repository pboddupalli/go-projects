package strings

import "strconv"

type TokenType int

const (
  SimpleString TokenType = iota
  RepeatedString
)

type Token struct {
  tokenType TokenType
  count int
  s string
}

func decodeString(s string) string {
  return decode(s, 1)
}

func decode(s string, count int) string {
  splits := parse(s)
  var tmp string
  for _, split := range splits {
    if split.tokenType == SimpleString {
      tmp = tmp + split.s
    } else {
      tmp = tmp + decode(split.s, split.count)
    }
  }
  result := tmp
  for count > 1 {
    result = result + tmp
    count--
  }
  return result
}

func parse(s string) []Token {
  result := []Token{}
  var num int
  var val string

  for i := 0; i < len(s); {
    if isDigit(s[i]) {
      num, i = parseNumber(s, i)
      val, i = parseRepeatedString(s, i)
      result = append(result, Token{tokenType: RepeatedString, count: num, s: val})
    } else {
      val, i = parseSimpleString(s, i)
      result = append(result, Token{tokenType: SimpleString, s: val})
    }
  }
  return result
}

func isDigit(b byte) bool {
  return b >= '0' && b <= '9'
}

func isAlpha(b byte) bool {
  return b >= 'a' && b <= 'z'
}

func parseNumber(s string, i int) (count int, next int) {
  start := i
  for i < len(s) && isDigit(s[i]) {
    i++
  }
  num, _ := strconv.Atoi(s[start:i])
  return num, i
}

func parseSimpleString(s string, i int) (string, int) {
  start := i
  for i < len(s) && isAlpha(s[i]) {
    i++
  }
  return s[start:i], i
}

func parseRepeatedString(s string, i int) (string, int) {
  i++
  open, start := 1, i
  for ; i < len(s) && open > 0; i++ {
    if s[i] == '[' {
      open++
    } else if s[i] == ']' {
      open--
    }
  }
  return s[start:i-1], i
}
