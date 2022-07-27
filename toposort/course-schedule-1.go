package toposort

type node struct {
  prereqs []int
  seen int
}

func canFinish(n int, prerequisites [][]int) bool {

  graph := make([]node, n)

  for _, req := range prerequisites {
    graph[req[0]].prereqs = append(graph[req[0]].prereqs, req[1])
  }

  for i := range graph {
    if graph[i].seen == 0 {
      retval := dfs(graph, i)
      if retval == false {
        return false
      }
    }
  }
  return true
}

func dfs(graph []node, i int) bool {
  if graph[i].seen == -1 {
    // cycle
    return false
  }
  if graph[i].seen == 1 {
    return true
  }
  graph[i].seen = -1
  for _, req := range graph[i].prereqs {
    retval := dfs(graph, req)
    if retval == false {
      return false
    }
  }
  graph[i].seen = 1
  return true
}
