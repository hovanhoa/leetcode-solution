func dfs(seen []bool, rooms[][]int, node int) {
    seen[node] = true
    for _, v := range rooms[node] {
        if !seen[v] {
            dfs(seen, rooms, v)
        }
    }
}

func canVisitAllRooms(rooms [][]int) bool {
    seen := make([]bool, len(rooms))
    dfs(seen, rooms, 0)

    for _, v := range seen {
        if !v {
            return false
        }
    }

    return true
}