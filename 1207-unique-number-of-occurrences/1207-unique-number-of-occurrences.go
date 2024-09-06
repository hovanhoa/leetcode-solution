func uniqueOccurrences(arr []int) bool {
    occurencesMap := make(map[int]int)
    for _, v := range arr {
        occurencesMap[v]++
    }

    uniqueMap := make(map[int]bool)
    for _, v := range occurencesMap {
        if !uniqueMap[v] {
            uniqueMap[v] = true
        } else {
            return false
        }
    }

    return true
}