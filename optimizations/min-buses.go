package optimizations

import "sort"

func numBusesToDestination(routes [][]int, source, target int) int {
  if source == target {
    return 0
  }

  seen := make([]bool, len(routes))
  targets := make(map[int]bool)

  graph := make([][]int, len(routes))
  for i := 0; i < len(routes); i++ {
    graph[i] = make([]int, 0)
    sort.Ints(routes[i])
  }

  // step 1: create the graph, for each bus (route), the list of buses with which it shares a stop
  for i := 0; i < len(routes); i++ {
    for j := i + 1; j < len(routes); j++ {
      if intersects(routes[i], routes[j]) {
        graph[i] = append(graph[i], j)
        graph[j] = append(graph[j], i)
      }
    }
  }

  // step 2:
  queue := make([]int, 0)
  for i := 0; i < len(routes); i++ {
    if binSrch(routes[i], source)!= -1 {
      seen[i] = true
      queue = append(queue, i)
    }
    if binSrch(routes[i], target) != -1 {
      targets[i] = true
    }
  }

  // step 3: loop over elements in the queue
  steps := 0
  for len(queue) > 0 {
    currentBatch := queue
    steps++
    queue = queue[len(queue):]
    for _, bus := range currentBatch {
      _, ok := targets[bus]
      if ok {
        return steps
      }
      for _, child := range graph[bus] {
        if !seen[child] {
          seen[child] = true
          queue = append(queue, child)
        }
      }
    } // target of inner for loop

  } // target of looping over q elements
  return -1
} // end of function

func intersects(list1 []int, list2 []int) bool {
  for i := 0; i < len(list1); i++ {
    if binSrch(list2, list1[i]) != -1 {
      return true
    }
  }
  return false
}

func binSrch(nums []int, target int) int {
  i, j := 0, len(nums)
  for i < j {
    mid := i + (j - i)/2
    if nums[mid] == target {
      return mid
    } else if target > nums[mid] {
      i = mid + 1
    } else {
      j = mid
    }
  }
  return -1
}
