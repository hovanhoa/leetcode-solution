func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    w1Map := make(map[byte]int)
    w2Map := make(map[byte]int)

    for i := range word1 {
        w1Map[word1[i]]++
        w2Map[word2[i]]++
    }

    valueW2Map := make(map[int]int)
    for _, v := range w2Map {
        valueW2Map[v]++
    }

    for _, v := range w1Map {
        if valueW2Map[v] == 0 {
            return false
        }

        valueW2Map[v]--
    }

    keyW2Map := make(map[byte]bool)
    for k := range w2Map {
        keyW2Map[k] = true
    }

    for k := range w1Map {
        if !keyW2Map[k] {
            return false
        }
    }

    return true
}