package trees

func distributeCoins(root *TreeNode) int {
  cost, _ := findCost(root)
  return cost
}

func findCost(root *TreeNode) (int, int) {
  if root == nil {
    return 0,0
  }
  lCost, lExcess := findCost(root.Left)
  rCost, rExcess := findCost(root.Right)
  cost := lCost + rCost + abs(lExcess) + abs(rExcess)
  excess := lExcess + rExcess + root.Val - 1
  return cost, excess
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

