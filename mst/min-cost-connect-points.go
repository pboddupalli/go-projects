package mst

import (
  "sort"
)

type edge struct {
  node1, node2 int
  distance int
}

func minCostConnectPoints(points [][]int) int {
  edges := make([]edge, 0)
  for i := 0; i < len(points); i++ {
    for j := i + 1; j < len(points); j++ {
      edges = append(edges, edge{node1: i, node2: j,
        distance: abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1])})
    }
  }
  sort.Slice(edges, func(i, j int) bool { return edges[i].distance < edges[j].distance })

  parent := make([]int, len(points))
  for i := 0; i < len(points); i++ {
    parent[i] = i
  }

  cost, count := 0, 0
  for i := 0; count < len(points)-1; i++ {

    p1 := FindParent(parent, edges[i].node1)
    p2 := FindParent(parent, edges[i].node2)
    if p1 == p2 {
      continue
    }
    count++
    parent[p1] = p2
    cost += edges[i].distance
  }
  return cost
}

func FindParent(parent []int, i int) int {
  if parent[i] != i {
    parent[i] = FindParent(parent, parent[i])
  }
  return parent[i]
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}
