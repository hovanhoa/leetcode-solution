func uniqueOccurrences(arr []int) bool {
    m := map[int]int{}
    for _, v := range arr {
        m[v] += 1
    }

    occ := map[int]bool{}
    for _, v := range m {
        if occ[v] {
            return false
        }

        occ[v] = true
    }

    return true
}
