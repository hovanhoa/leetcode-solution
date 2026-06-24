func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    mWord1, mWord2 := map[byte]int{}, map[byte]int{}
    for i := 0; i < len(word1); i++ {
        mWord1[word1[i]]++
        mWord2[word2[i]]++
    }

    // compare value
    vWord1 := map[int]int{}
    kWord1 := map[byte]bool{}
    for k, v := range mWord1 {
        vWord1[v]++
        kWord1[k] = true
    }

    for k, v := range mWord2 {
        if vWord1[v] == 0 {
            return false
        }

        if !kWord1[k] {
            return false
        }

        vWord1[v]--
    }

    return true
}