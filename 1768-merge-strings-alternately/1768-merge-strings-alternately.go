func mergeAlternately(word1 string, word2 string) string {
    i, j := 0, 0
    res := ""
    for ;i < len(word1) || j < len(word2); i, j = i+1, j+1 {
        if i < len(word1) {
            res += string(word1[i])
        }

        if j < len(word2) {
            res += string(word2[j])
        }
    }

    return res
}