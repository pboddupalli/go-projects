package intervals

import (
  "fmt"
  "testing"
)

func TestMergeIntervals(t *testing.T) {
  var intervals [][]int

  intervals = [][]int{{1,3},{2,6},{8,10},{15,18}}

  fmt.Println(mergeIntervalsVer2(intervals))
}
