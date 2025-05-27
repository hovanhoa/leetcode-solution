func uniqueOccurrences(arr []int) bool {
    occurrences := map[int]int{}
    for _, v := range arr {
        occurrences[v] += 1
    }

    m := map[int]bool{}
    for _, v := range occurrences {
        if m[v] {
            return false
        }

        m[v] = true
    }

    return true
}