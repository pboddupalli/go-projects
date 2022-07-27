package mst

import "math"

func minCostConnectPointsVer2(points [][]int) int {
  totalCost := 0
  current := points[0]

  costs := make([]int, len(points) - 1)
  points = points[1:]
  for i := 0; i < len(points); i++ {
    costs[i] = math.MaxInt32
  }

  for {
    nextIndex := -1
    minCost := math.MaxInt32
    for i := 0; i < len(points); i++ {
      costs[i] = min(costs[i], getDist(current, points[i]))
      if costs[i] < minCost {
        nextIndex = i
        minCost = costs[i]
      }
    }

    if nextIndex == -1 {
      break
    }

    totalCost += costs[nextIndex]
    current = points[nextIndex]

    points[nextIndex], points[len(points) - 1] = points[len(points) - 1], points[nextIndex]
    costs[nextIndex], costs[len(costs) - 1] = costs[len(costs) - 1], costs[nextIndex]

    points = points[:len(points) - 1]
    costs = costs[:len(costs) - 1]
  }
  return totalCost
}

func getDist(p1, p2 []int) int {
  return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])
}

func min(i, j int) int {
  if i < j {
    return i
  }
  return j
}
