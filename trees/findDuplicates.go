package trees

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
  count := make(map[string]int)
  var res []*TreeNode
  dfs(root, count, &res)
  return res
}

func dfs(root *TreeNode, count map[string]int, res *[]*TreeNode) string {
  if root == nil {
    return "nil"
  }
  l := dfs(root.Left, count, res)
  r := dfs(root.Right, count, res)
  path := strconv.Itoa(root.Val) + "#" + l + "#" + r
  count[path]++
  if count[path] == 2 {
    *res = append(*res, root)
  }
  return path
}
