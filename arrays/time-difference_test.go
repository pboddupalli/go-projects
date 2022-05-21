package arrays

import (
  "fmt"
  "testing"
)

func TestFindMinTimeDifference(t *testing.T) {
  fmt.Println(findMinDifference([]string{"23:59","00:00"}))
  fmt.Println(findMinDifference([]string{"00:00","23:59","00:00"}))
}
