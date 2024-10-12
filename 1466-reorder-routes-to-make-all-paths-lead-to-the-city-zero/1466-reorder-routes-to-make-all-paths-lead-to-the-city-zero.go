func minReorder(n int, connections [][]int) int {
    graph := make([][]int, n)
    for i := 0; i < len(connections); i++ {
        f, s := connections[i][0], connections[i][1]
        graph[f] = append(graph[f], s)
        graph[s] = append(graph[s], -f)
    }

    fmt.Println(graph)

    visited := make([]bool, n)
    res := 0
    var dfs func(n int) 
    dfs = func(num int) {
        visited[num] = true
        for i := 0; i < len(graph[num]); i++ {
            if !visited[abs(graph[num][i])] {
                dfs(abs(graph[num][i]))
                if graph[num][i] > 0 {
                    res++
                }
            }
        }
    }

    for i := 0; i < len(graph); i++ {
        if !visited[i] {
            dfs(i)
        }
    }

    return res
}

func abs(n int) int {
    if n >= 0 {
        return n
    }

    return -n
}