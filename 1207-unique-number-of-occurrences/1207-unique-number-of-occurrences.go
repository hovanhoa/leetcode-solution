func uniqueOccurrences(arr []int) bool {
    occurrences := make(map[int]int)
    for _, val := range arr {
        occurrences[val]++
    }

    exists := make(map[int]bool)
    for _, count := range occurrences {
        if exists[count] {
            return false
        }

        exists[count] = true
    }

    return true
}