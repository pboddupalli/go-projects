package intervals

func insert(intervals [][]int, new []int) [][]int {
  result := [][]int{}
  var i int
  for i = 0; i < len(intervals) && intervals[i][1] < new[0]; i++ {
    result = append(result, intervals[i])
  }
  if i < len(intervals) && intervals[i][0] < new[0] {
    new[0] = intervals[i][0]
  }
  for ; i < len(intervals) && intervals[i][1] <= new[1]; i++ {

  }
  if i < len(intervals) && intervals[i][0] <= new[1] && intervals[i][1] > new[1] {
    new[1] = intervals[i][1]
    i++
  }
  result = append(result, new)
  result = append(result, intervals[i:]...)
  return result
}
