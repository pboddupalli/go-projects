package hard

import "sort"

func maxEnvelopes(envelopes [][]int) int {
  sort.Slice(envelopes, func (i, j int) bool {
    if envelopes[i][0] == envelopes[j][0] { // if heights are equal, compare weights
      return envelopes[i][1] > envelopes[j][1]
    }
    return envelopes[i][0] > envelopes[j][0]
  })

  result := make([]int, len(envelopes))
  result[0] = 1
  max := 1
  for i := 1; i < len(envelopes); i++ {
    result[i] = 1
    for j := i-1; j >= 0; j-- {
      if envelopes[i][1] < envelopes[j][1] && envelopes[i][0] < envelopes[j][0] && result[j] + 1 > result[i] {
        result[i] = result[j] + 1
      }
    }
    if result[i] > max {
      max = result[i]
    }
  }
  return max
}
