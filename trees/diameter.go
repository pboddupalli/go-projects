package trees

import "fmt"

func diameterOfBinaryTree(root *TreeNode) int {
  max, _ := diameter(root)
  return max
}

func diameter(root *TreeNode) (int, int) {
  if root == nil {
    return 0, -1
  }
  leftDia, leftMaxPath := diameter(root.Left)
  rightDia, rightMaxPath := diameter(root.Right)
  dia := max(leftDia, rightDia, leftMaxPath + rightMaxPath + 2)
  fmt.Println(root.Val, dia)
  return dia, max(leftMaxPath, rightMaxPath) + 1
}

func max(nums ...int) int {
  result := nums[0]
  for _, num := range nums[1:] {
    if num > result {
      result = num
    }
  }
  return result
}
