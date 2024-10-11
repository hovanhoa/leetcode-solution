func dfs(visited map[int]bool, isConnected [][]int, startCity int) {
    for i := range isConnected[startCity] {
        if isConnected[startCity][i] == 1 && !visited[i] {
            visited[i] = true
            dfs(visited, isConnected, i)
        }
    }
}

func findCircleNum(isConnected [][]int) int {
    visited := map[int]bool{}
    ans := 0
    for i := range isConnected {
        if !visited[i] {
            dfs(visited, isConnected, i)
            ans += 1
        }
    }

    return ans
}