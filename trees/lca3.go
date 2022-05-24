package trees

import "fmt"

func lca3(p, q *Node) *Node {
  nodeMap := make(map[*Node]bool)
  for p != nil || q != nil {
    if p != nil {
      _, ok := nodeMap[p]
      if ok {
        return p
      }
      fmt.Println("p inserting...", p.Val)
      nodeMap[p] = true
      p = p.Parent
    }
    if q != nil {
      fmt.Println(q.Val)
      _, ok := nodeMap[q]
      if ok {
        return q
      }
      fmt.Println("q inserting...", q.Val)
      nodeMap[q] = true
      q = q.Parent
    }
  }
  return nil
}
