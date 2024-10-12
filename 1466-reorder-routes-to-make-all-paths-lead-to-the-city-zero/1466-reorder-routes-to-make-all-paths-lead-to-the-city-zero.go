func minReorder(n int, connections [][]int) int {
    sort.Slice(connections, func(i, j int) bool {
        return connections[i][1] < connections[j][1]
    })

    ans := 0
    visited := map[int]bool{0: true}
    for i := range connections {
        if !visited[connections[i][1]] {
            visited[connections[i][1]] = true
            ans += 1
        } else {
            visited[connections[i][0]] = true
        }
    }

    return ans
}