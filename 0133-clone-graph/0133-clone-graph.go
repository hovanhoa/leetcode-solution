/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func dfs(node *Node, visited map[*Node]*Node) *Node {
    if node == nil {
        return nil
    }

    if val, ok := visited[node]; ok {
        return val
    }

    visited[node] = &Node{Val: node.Val}
    for _, nei := range node.Neighbors {
        visited[node].Neighbors = append(visited[node].Neighbors, dfs(nei, visited))
    }

    return visited[node]
}

func cloneGraph(node *Node) *Node {
    visited := map[*Node]*Node{}
    return dfs(node, visited)
}