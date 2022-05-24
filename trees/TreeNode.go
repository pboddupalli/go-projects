package trees

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func BuildBinaryTree(elems []interface{}) *TreeNode {
  if len(elems) == 0 {
    return nil
  }
  root := &TreeNode{Val: elems[0].(int)}
  q := []*TreeNode{root}
  for i := 1; len(q) != 0 && i < len(elems); i++ {
    current := q[0]
    q = q[1:]
    if elems[i] != nil {
      current.Left = &TreeNode{Val: elems[i].(int)}
      q = append(q, current.Left)
    }
    i++
    if i == len(elems) {
      break;
    }
    if elems[i] != nil {
      current.Right = &TreeNode{Val: elems[i].(int)}
      q = append(q, current.Right)
    }
  }
  return root
}