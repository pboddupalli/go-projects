package trees

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  // at root level we are bound to get non-nil paths since p and q are assumed to be valid pointers
  leftPath := getPath(root, p)
  rightPath := getPath(root, q)

  i, j := len(leftPath)-1, len(rightPath)-1
  var result *TreeNode
  for i >= 0 && j >= 0 && leftPath[i] == rightPath[j] {
    result = leftPath[i]
    i--; j--
  }
  return result
}

func getPath(root, p *TreeNode) []*TreeNode {
  if root == nil {
    return nil
  }
  if root == p {
    return []*TreeNode{p}
  }
  leftResult := getPath(root.Left, p)
  if leftResult != nil {
    return append(leftResult, root)
  }
  rightResult := getPath(root.Right, p)
  if rightResult != nil {
    return append(rightResult, root)
  }
  return nil
}
