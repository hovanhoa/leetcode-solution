func canVisitAllRooms(rooms [][]int) bool {
    visited := make(map[int]bool)
    keys := []int{0}
    for len(keys) > 0 {
        newKeys := []int{}
        for _, k := range keys {
            if !visited[k] {
                newKeys = append(newKeys, rooms[k]...)
            }

            visited[k] = true
        }

        keys = newKeys
    }

    return len(visited) == len(rooms)
}