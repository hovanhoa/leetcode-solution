func uniqueOccurrences(arr []int) bool {
    occ := make(map[int]int, 0)
    for _, v := range arr {
        occ[v]++
    }

    exist := make(map[int]bool)
    for _, v := range occ {
        if exist[v] {
            return false
        }

        exist[v] = true
    }

    return true
}